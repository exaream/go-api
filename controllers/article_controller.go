package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/exaream/go-api/apperrors"
	"github.com/exaream/go-api/controllers/services"
	"github.com/exaream/go-api/models"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) GetArticleListHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "page must be number")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	articleList, err := c.service.GetArticleList(page)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
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

func (c *ArticleController) GetArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpID := r.PathValue("id")
	if !regexp.MustCompile(`\d+`).MatchString(tmpID) {
		err := apperrors.BadParam.Wrap(nil, "id must be number")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	id, err := strconv.Atoi(tmpID)
	if err != nil {
		err := apperrors.BadParam.Wrap(nil, "failed to convert id")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	article, err := c.service.GetArticleDetail(id)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode request body")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	article, err := c.service.PostArticle(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	article, err := c.service.PostNice(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}
