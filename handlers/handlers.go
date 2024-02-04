package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/exaream/go-api/models"
)

func GetArticleListHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if err != nil {
		http.Error(w, "invalid page", http.StatusBadRequest)
		return
	}

	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}
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
	id := r.PathValue("id")
	if !regexp.MustCompile(`\d+`).MatchString(id) {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article.NiceNum++
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
