package services

import (
	"github.com/exaream/go-api/internal/apperrors"
	"github.com/exaream/go-api/internal/models"
	"github.com/exaream/go-api/internal/repositories"
)

func (s *AppService) PostComment(comment *models.Comment) (*models.Comment, error) {
	comment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return nil, apperrors.FailedToInsert.Wrap(err, "failed to insert comment")
	}

	return comment, nil
}
