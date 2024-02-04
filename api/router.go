package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/controllers"
	"github.com/exaream/go-api/services"
)

func NewRouter(db *sql.DB, logger *slog.Logger) *http.ServeMux {
	srv := services.NewAppService(db, logger)
	articleCtrl := controllers.NewArticleController(srv)
	commentCtrl := controllers.NewCommentController(srv)
	mux := http.NewServeMux()

	// Article
	mux.HandleFunc("GET /article/list", articleCtrl.GetArticleListHandler)
	mux.HandleFunc("GET /article/{id}", articleCtrl.GetArticleDetailHandler)
	mux.HandleFunc("POST /article", articleCtrl.PostArticleHandler)
	mux.HandleFunc("POST /article/nice", articleCtrl.PostNiceHandler)

	// Comment
	mux.HandleFunc("POST /comment", commentCtrl.PostCommentHandler)

	return mux
}
