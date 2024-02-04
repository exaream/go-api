package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/exaream/go-api/controllers/services"
	"github.com/exaream/go-api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqComment *models.Comment
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostComment(reqComment)
	if err != nil {
		http.Error(w, "failed to post comment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
