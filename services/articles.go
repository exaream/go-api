package services

import (
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func GetArticleList(page int) ([]*models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil

}

func GetArticleDetail(articleID int) (*models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return nil, err
	}

	comments, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return nil, err
	}

	article.CommentList = comments

	return article, nil
}

func PostArticle(article *models.Article) (*models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return repositories.InsertArticle(db, article)
}

func PostNice(article *models.Article) (*models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	newArticle, err := repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return nil, err
	}

	return newArticle, nil
}
