package services

import (
	"github.com/exaream/go-api/apperrors"
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func (s *AppService) PostComment(comment *models.Comment) (*models.Comment, error) {
	comment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return nil, apperrors.InsertDataFailed.Wrap(err, "failed to insert comment")
	}

	return comment, nil
}
