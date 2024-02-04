package handlers

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func GetArticleListHandler(w http.ResponseWriter, r *http.Request) {
	page, err := getPage(r)
	if err != nil {
		http.Error(w, "invalid page", http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("get article list: page=%d\n", page))
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

	io.WriteString(w, fmt.Sprintf("get article detail: id=%s\n", id))
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "post article\n")
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "post nice\n")
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "post comment\n")
}
