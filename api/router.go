package api

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/api/middlewares"
	"github.com/exaream/go-api/controllers"
	"github.com/exaream/go-api/services"
)

func NewRouter(ctx context.Context, logger *slog.Logger, db *sql.DB) *http.ServeMux {
	srv := services.NewAppService(logger, db)
	articleCtrl := controllers.NewArticleController(ctx, logger, srv)
	commentCtrl := controllers.NewCommentController(ctx, logger, srv)
	mux := http.NewServeMux()
	middlewareList := []func(context.Context, *slog.Logger, http.HandlerFunc) http.HandlerFunc{middlewares.LoggingMiddleware}

	mux.HandleFunc("GET /article/list", middlewares.Apply(ctx, logger, articleCtrl.ListArticle, middlewareList))
	mux.HandleFunc("GET /article/{id}", middlewares.Apply(ctx, logger, articleCtrl.GetArticle, middlewareList))
	mux.HandleFunc("POST /article", middlewares.Apply(ctx, logger, articleCtrl.PostArticle, middlewareList))
	mux.HandleFunc("POST /article/nice", middlewares.Apply(ctx, logger, articleCtrl.PostNice, middlewareList))
	mux.HandleFunc("POST /comment", middlewares.Apply(ctx, logger, commentCtrl.PostComment, middlewareList))

	return mux
}
