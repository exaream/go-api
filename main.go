package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/exaream/go-api/handlers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /article/list", handlers.GetArticleListHandler)
	mux.HandleFunc("GET /article/{id}", handlers.GetArticleDetailHandler)
	mux.HandleFunc("POST /article", handlers.PostArticleHandler)
	mux.HandleFunc("POST /article/nice", handlers.PostNiceHandler)
	mux.HandleFunc("POST /comment", handlers.PostCommentHandler)

	port := os.Getenv("HTTP_PORT")
	logger.InfoContext(ctx, "starting server", slog.String("port", port))

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		logger.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}
