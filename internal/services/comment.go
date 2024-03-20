package services

import (
	"database/sql"
	"log/slog"

	"github.com/exaream/go-api/internal/apperrors"
	"github.com/exaream/go-api/internal/models"
	"github.com/exaream/go-api/internal/repositories"
)

type CommentService struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewCommentService(logger *slog.Logger, db *sql.DB) *CommentService {
	return &CommentService{
		logger: logger,
		db:     db,
	}
}

func (s *CommentService) Post(comment *models.Comment) (*models.Comment, error) {
	comment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return nil, apperrors.FailedToInsert.Wrap(err, "failed to insert comment")
	}

	return comment, nil
}
