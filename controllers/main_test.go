package controllers_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/exaream/go-api/controllers"
	"github.com/exaream/go-api/controllers/testdata"
)

var articleCtrl *controllers.ArticleController

func TestMain(m *testing.M) {
	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	ctx := context.Background()

	srv := testdata.NewServiceMock()
	articleCtrl = controllers.NewArticleController(ctx, logger, srv)

	m.Run()
}
