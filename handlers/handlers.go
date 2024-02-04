package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/services"
)

func GetArticleListHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if err != nil {
		http.Error(w, "invalid page", http.StatusBadRequest)
		return
	}

	log.Println(page)

	articleList, err := services.GetArticleList(page)
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

func GetArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
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

	article, err := services.GetArticleDetail(id)
	if err != nil {
		http.Error(w, "failed to get article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article, err := services.PostArticle(&reqArticle)
	if err != nil {
		http.Error(w, "failed to post article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article, err := services.PostNice(&reqArticle)
	if err != nil {
		http.Error(w, "failed to post nice", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	comment, err := services.PostComment(&reqComment)
	if err != nil {
		http.Error(w, "failed to post comment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
