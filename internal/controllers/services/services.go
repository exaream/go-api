package services

import (
	"github.com/exaream/go-api/internal/models"
)

type ArticleServicer interface {
	ListArticle(page int) ([]*models.Article, error)
	GetArticle(id int) (*models.Article, error)
	PostArticle(article *models.Article) (*models.Article, error)
	PostNice(article *models.Article) (*models.Article, error)
}

type CommentServicer interface {
	PostComment(comment *models.Comment) (*models.Comment, error)
}
