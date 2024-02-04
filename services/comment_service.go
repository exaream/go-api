package services

import (
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func (s *AppService) PostComment(comment *models.Comment) (*models.Comment, error) {
	return repositories.InsertComment(s.db, comment)
}
