package services

import (
	"database/sql"
	"log/slog"
)

type ArticleService struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewArticleService(logger *slog.Logger, db *sql.DB) *ArticleService {
	return &ArticleService{
		logger: logger,
		db:     db,
	}
}

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
