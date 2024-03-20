package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/exaream/go-api/internal/config"
	"github.com/exaream/go-api/internal/database"
	"github.com/exaream/go-api/internal/router"
)

func main() {
	var dotenvPath = flag.String("env", ".env", "file path of .env")
	flag.Parse()

	opt := &slog.HandlerOptions{Level: slog.LevelDebug} // AddSource: true
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	ctx := context.Background()

	if err := config.Load(*dotenvPath); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	cfg, err := config.Get()
	if err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	db, err := database.Connect("mysql", &cfg.DB)
	if err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	logger.InfoContext(ctx, "starting server")

	handler := router.NewHandler(ctx, cfg, logger, db)
	if err := http.ListenAndServe(cfg.HTTPPort, handler); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}
