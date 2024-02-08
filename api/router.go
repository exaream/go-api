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
	mux := http.NewServeMux()
	handleArticle(mux, srv)
	handleComment(mux, srv)

	return mux
}

func handleArticle(mux *http.ServeMux, srv *services.AppService) {
	ctrl := controllers.NewArticleController(srv)
	mux.HandleFunc("GET /article/list", middlewares.LoggingMiddleware(ctrl.GetArticleListHandler))
	mux.HandleFunc("GET /article/{id}", middlewares.LoggingMiddleware(ctrl.GetArticleDetailHandler))
	mux.HandleFunc("POST /article", middlewares.LoggingMiddleware(ctrl.PostArticleHandler))
	mux.HandleFunc("POST /article/nice", middlewares.LoggingMiddleware(ctrl.PostNiceHandler))
}

func handleComment(mux *http.ServeMux, srv *services.AppService) {
	ctrl := controllers.NewCommentController(srv)
	mux.HandleFunc("POST /comment", middlewares.LoggingMiddleware(ctrl.PostCommentHandler))
}
