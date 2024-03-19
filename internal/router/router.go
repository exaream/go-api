package router

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/internal/controllers"
	"github.com/exaream/go-api/internal/middlewares"
	"github.com/exaream/go-api/internal/services"
)

func NewHandler(ctx context.Context, logger *slog.Logger, db *sql.DB) *http.ServeMux {
	srv := services.NewAppService(logger, db)
	articleCtrl := controllers.NewArticleController(ctx, logger, srv)
	commentCtrl := controllers.NewCommentController(ctx, logger, srv)
	middleware := middlewares.NewMiddleware(logger)
	middlewareList := []func(*slog.Logger, http.HandlerFunc) http.HandlerFunc{middlewares.Logging}

	// If you want to use both URLs with and without a trailing slash,
	// please make sure to set a slash at the end of the URL.
	mux := http.NewServeMux()
	mux.Handle("GET /article/list/{$}", middleware.Chain(articleCtrl.ListArticle, middlewareList))
	mux.Handle("GET /article/{id}/{$}", middleware.Chain(articleCtrl.GetArticle, middlewareList))
	mux.Handle("POST /article/{$}", middleware.Chain(articleCtrl.PostArticle, middlewareList))
	mux.Handle("POST /article/nice/{$}", middleware.Chain(articleCtrl.PostNice, middlewareList))
	mux.Handle("POST /comment/{$}", middleware.Chain(commentCtrl.PostComment, middlewareList))

	return mux
}
