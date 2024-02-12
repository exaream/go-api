package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/exaream/go-api/api"
)

func main() {
	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	ctx := context.Background()

	db, err := getDB()
	if err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	router := api.NewRouter(db, logger)

	logger.InfoContext(ctx, "starting server")
	if err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}
