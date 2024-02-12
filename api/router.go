package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/api/middlewares"
	"github.com/exaream/go-api/controllers"
	"github.com/exaream/go-api/services"
)

func NewRouter(db *sql.DB, logger *slog.Logger) *http.ServeMux {
	srv := services.NewAppService(db, logger)
	articleCtrl := controllers.NewArticleController(srv)
	commentCtrl := controllers.NewCommentController(srv)
	mux := http.NewServeMux()
	middlewareList := []func(http.HandlerFunc) http.HandlerFunc{middlewares.Logging}

	mux.HandleFunc("GET /article/list", middlewares.Apply(articleCtrl.GetArticleListHandler, middlewareList))
	mux.HandleFunc("GET /article/{id}", middlewares.Apply(articleCtrl.GetArticleDetailHandler, middlewareList))
	mux.HandleFunc("POST /article", middlewares.Apply(articleCtrl.PostArticleHandler, middlewareList))
	mux.HandleFunc("POST /article/nice", middlewares.Apply(articleCtrl.PostNiceHandler, middlewareList))
	mux.HandleFunc("POST /comment", middlewares.Apply(commentCtrl.PostCommentHandler, middlewareList))

	return mux
}
