package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/exaream/go-api/internal/database"
	"github.com/exaream/go-api/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	var dotenvPath = flag.String("env", ".env", "file path of .env")
	flag.Parse()

	opt := &slog.HandlerOptions{Level: slog.LevelDebug} // AddSource: true
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	ctx := context.Background()

	if err := godotenv.Load(*dotenvPath); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	db, err := database.Connect("mysql", "")
	if err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}

	logger.InfoContext(ctx, "starting server")

	handler := router.NewHandler(ctx, logger, db)
	if err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), handler); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}
