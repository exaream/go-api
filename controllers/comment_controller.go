package controllers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/apperrors"
	"github.com/exaream/go-api/controllers/services"
	"github.com/exaream/go-api/models"
)

type CommentController struct {
	ctx     context.Context
	logger  *slog.Logger
	service services.CommentServicer
}

func NewCommentController(ctx context.Context, logger *slog.Logger, service services.CommentServicer) *CommentController {
	return &CommentController{
		ctx:     ctx,
		logger:  logger,
		service: service,
	}
}

func (c *CommentController) PostComment(w http.ResponseWriter, r *http.Request) {
	var reqComment *models.Comment
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		err = apperrors.FailedToDecodeReq.Wrap(err, "failed to decode request body")
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
