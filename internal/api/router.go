package api

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/internal/api/middlewares"
	"github.com/exaream/go-api/internal/controllers"
	"github.com/exaream/go-api/internal/services"
)

func NewRouter(ctx context.Context, logger *slog.Logger, db *sql.DB) *http.ServeMux {
	srv := services.NewAppService(logger, db)
	articleCtrl := controllers.NewArticleController(ctx, logger, srv)
	commentCtrl := controllers.NewCommentController(ctx, logger, srv)
	middleware := middlewares.NewMiddleware(logger)
	middlewareList := []func(*slog.Logger, http.HandlerFunc) http.HandlerFunc{middlewares.LoggingMiddleware}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /article/list", middleware.Apply(articleCtrl.ListArticle, middlewareList))
	mux.HandleFunc("GET /article/{id}", middleware.Apply(articleCtrl.GetArticle, middlewareList))
	mux.HandleFunc("POST /article", middleware.Apply(articleCtrl.PostArticle, middlewareList))
	mux.HandleFunc("POST /article/nice", middleware.Apply(articleCtrl.PostNice, middlewareList))
	mux.HandleFunc("POST /comment", middleware.Apply(commentCtrl.PostComment, middlewareList))

	return mux
}
