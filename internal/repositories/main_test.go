package repositories_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/exaream/go-api/internal/config"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var cfg *config.Config
	var err error

	cfg, testDB, err = Setup()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to set up test: %w", err))
	}

	defer testDB.Close()

	m.Run()

	if err := Teardown(cfg); err != nil {
		log.Fatalln(fmt.Errorf("failed to tear down test: %w", err))
	}
}
