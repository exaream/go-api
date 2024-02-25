package services

import (
	"database/sql"
	"log/slog"
)

type AppService struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewAppService(logger *slog.Logger, db *sql.DB) *AppService {
	return &AppService{
		logger: logger,
		db:     db,
	}
}
