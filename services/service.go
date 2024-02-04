package services

import (
	"database/sql"
	"log/slog"
)

type AppService struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewAppService(db *sql.DB, logger *slog.Logger) *AppService {
	return &AppService{db: db, logger: logger}
}
