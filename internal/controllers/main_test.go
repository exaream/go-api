package controllers_test

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/exaream/go-api/internal/controllers"
	"github.com/exaream/go-api/internal/controllers/testdata"
)

var articleCtrl *controllers.ArticleController

func TestMain(m *testing.M) {
	cfg, _, err := Setup()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to set up test: %w", err))
	}

	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	ctx := context.Background()

	srv := testdata.NewArticleServiceMock()
	articleCtrl = controllers.NewArticleController(ctx, logger, srv)

	m.Run()

	if err := Teardown(cfg); err != nil {
		log.Fatalln(fmt.Errorf("failed to tear down test: %w", err))
	}
}
