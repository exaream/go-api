package services_test

import (
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/exaream/go-api/internal/config"
	"github.com/exaream/go-api/internal/database"
	"github.com/exaream/go-api/internal/services"
)

const dotenvPath = "../../.env"

var articleService *services.ArticleService

func TestMain(m *testing.M) {
	config.Load(dotenvPath)
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	testDB, err := database.Connect("mysql", &cfg.TestDB)
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()

	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))

	articleService = services.NewArticleService(logger, testDB)

	m.Run()
}
