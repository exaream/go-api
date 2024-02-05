package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/exaream/go-api/api"
	_ "github.com/go-sql-driver/mysql"
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

func getDB() (*sql.DB, error) {
	src := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", src)
	if err != nil {
		return nil, err
	}

	return db, nil
}
