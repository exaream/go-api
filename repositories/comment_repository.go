package repositories

import (
	"database/sql"
	"time"

	"github.com/exaream/go-api/models"
)

func InsertComment(db *sql.DB, comment *models.Comment) (*models.Comment, error) {
	const query = `
		INSERT INTO comments (article_id, body, user_name, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?);
	`

	now := time.Now()
	res, err := db.Exec(query, comment.ArticleID, comment.Body, comment.UserName, now, now)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return models.NewComment(int(id), comment.Body, comment.UserName, now, now), nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]*models.Comment, error) {
	const query = `
		SELECT id, body, user_name, created_at, updated_at
		FROM comments
		WHERE article_id = ?
		ORDER BY created_at;
	`

	rows, err := db.Query(query, articleID)
	if err != nil {
		return nil, err
	}

	list := make([]*models.Comment, 0)
	for rows.Next() {
		var (
			comment   models.Comment
			createdAt sql.NullTime
			updatedAt sql.NullTime
		)

		err := rows.Scan(&comment.ID, &comment.Body, &comment.UserName, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		list = append(list, &comment)
	}

	return list, nil
}
