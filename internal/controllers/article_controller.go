package controllers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"

	"github.com/exaream/go-api/internal/apperrors"
	"github.com/exaream/go-api/internal/controllers/services"
	"github.com/exaream/go-api/internal/models"
)

type ArticleController struct {
	ctx     context.Context
	logger  *slog.Logger
	service services.ArticleServicer
}

func NewArticleController(ctx context.Context, logger *slog.Logger, service services.ArticleServicer) *ArticleController {
	return &ArticleController{
		ctx:     ctx,
		logger:  logger,
		service: service,
	}
}

func (c *ArticleController) ListArticle(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "page must be number")
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	articleList, err := c.service.ListArticle(page)
	if err != nil {
		c.logger.ErrorContext(c.ctx, "GetList", slog.Int("page", page))
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func getPage(r *http.Request) (int, error) {
	tmp, ok := r.URL.Query()["page"]
	if !ok || len(tmp) == 0 {
		return 1, nil
	}

	return strconv.Atoi(tmp[0])
}

func (c *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	tmpID := r.PathValue("id")
	if !regexp.MustCompile(`\d+`).MatchString(tmpID) {
		err := apperrors.BadParam.Wrap(nil, "id must be number")
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	id, err := strconv.Atoi(tmpID)
	if err != nil {
		err := apperrors.BadParam.Wrap(nil, "failed to convert id")
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	article, err := c.service.GetArticle(id)
	if err != nil {
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostArticle(w http.ResponseWriter, r *http.Request) {
	var reqArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		err = apperrors.FailedToDecodeReq.Wrap(err, "failed to decode request body")
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	article, err := c.service.PostArticle(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNice(w http.ResponseWriter, r *http.Request) {
	var reqArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	article, err := c.service.PostNice(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, c.logger, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}
