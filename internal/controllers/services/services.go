package services

import (
	"github.com/exaream/go-api/internal/models"
)

type ArticleServicer interface {
	List(page int) ([]*models.Article, error)
	GetByID(id int) (*models.Article, error)
	Post(article *models.Article) (*models.Article, error)
	PostNice(article *models.Article) (*models.Article, error)
}

type CommentServicer interface {
	Post(comment *models.Comment) (*models.Comment, error)
}
