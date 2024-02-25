package repositories

import (
	"database/sql"
	"time"

	"github.com/exaream/go-api/internal/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article *models.Article) (*models.Article, error) {
	const query = `
		INSERT INTO articles (title, body, user_name, nice_num, created_at, updated_at)
		VALUES (?, ?, ?, 0, ?, ?);
	`

	now := time.Now()
	res, err := db.Exec(query, article.Title, article.Body, article.UserName, now, now)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return models.NewArticle(int(id), article.Title, article.Body, article.UserName, 0, now, now), nil
}

func SelectArticleList(db *sql.DB, page int) ([]*models.Article, error) {
	const query = `
		SELECT id, title, body, user_name, nice_num, created_at, updated_at
		FROM articles
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(query, articleNumPerPage, articleNumPerPage*(page-1))
	if err != nil {
		return nil, err
	}

	list := make([]*models.Article, 0)
	for rows.Next() {
		var (
			article   models.Article
			createdAt sql.NullTime
			updatedAt sql.NullTime
		)

		err := rows.Scan(&article.ID, &article.Title, &article.Body, &article.UserName, &article.NiceNum, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		article.CreatedAt = createdAt.Time
		article.UpdatedAt = updatedAt.Time
		list = append(list, &article)
	}

	return list, nil
}

func SelectArticleDetail(db *sql.DB, id int) (*models.Article, error) {
	const query = `
		SELECT id, title, body, user_name, nice_num, created_at, updated_at
		FROM articles
		WHERE id = ?;
	`

	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var (
		article   models.Article
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)

	err := row.Scan(&article.ID, &article.Title, &article.Body, &article.UserName, &article.NiceNum, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	article.CreatedAt = createdAt.Time
	article.UpdatedAt = updatedAt.Time

	return &article, nil
}

func UpdateNiceNum(db *sql.DB, id int) (*models.Article, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	const querySelect = `
		SELECT id, title, body, user_name, nice_num
		FROM articles
		WHERE id = ?
		FOR UPDATE;
	`

	row := tx.QueryRow(querySelect, id)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return nil, err
	}

	var (
		article models.Article
	)

	err = row.Scan(&article.ID, &article.Title, &article.Body, &article.UserName, &article.NiceNum)
	if err != nil {
		return nil, err
	}

	const queryUpdate = `
		UPDATE articles
		SET nice_num = ?, updated_at = ?
		WHERE id = ?;
	`

	article.UpdatedAt = time.Now()
	article.NiceNum = article.NiceNum + 1
	_, err = tx.Exec(queryUpdate, article.NiceNum, article.UpdatedAt, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &article, nil
}
