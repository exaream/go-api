package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/exaream/go-api/apperrors"
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
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode request body")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	comment, err := c.service.PostComment(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
