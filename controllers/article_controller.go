package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		http.Error(w, "invalid page", http.StatusBadRequest)
		return
	}

	log.Println(page)

	articleList, err := c.service.GetArticleList(page)
	if err != nil {
		http.Error(w, "failed to get article list", http.StatusInternalServerError)
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
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(tmpID)
	if err != nil {
		http.Error(w, "failed to convert id", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleDetail(id)
	if err != nil {
		http.Error(w, "failed to get article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostArticle(reqArticle)
	if err != nil {
		http.Error(w, "failed to post article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostNice(reqArticle)
	if err != nil {
		http.Error(w, "failed to post nice", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
