package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/exaream/go-api/internal/api"
	"github.com/exaream/go-api/internal/database"
)

func main() {
	opt := &slog.HandlerOptions{Level: slog.LevelDebug} // AddSource: true
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	ctx := context.Background()

	db, err := database.Connect("mysql", "")
	if err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	logger.InfoContext(ctx, "starting server")

	router := api.NewRouter(ctx, logger, db)
	if err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}
