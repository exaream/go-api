package services

import (
	"github.com/exaream/go-api/models"
)

type ArticleServicer interface {
	GetArticleList(page int) ([]*models.Article, error)
	GetArticleDetail(id int) (*models.Article, error)
	PostArticle(article *models.Article) (*models.Article, error)
	PostNice(article *models.Article) (*models.Article, error)
}

type CommentServicer interface {
	PostComment(comment *models.Comment) (*models.Comment, error)
}
