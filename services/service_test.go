package services_test

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/exaream/go-api/database"
	"github.com/exaream/go-api/services"
)

const envPrefix = "TEST_"

var testDB *sql.DB
var articleService *services.AppService

func TestMain(m *testing.M) {
	var err error
	testDB, err = setup()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to setup test: %w", err))
	}

	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	articleService = services.NewAppService(logger, testDB)

	m.Run()

	if err := teardown(testDB); err != nil {
		log.Fatalln(fmt.Errorf("failed to teardown test: %w", err))
	}

}

func setup() (*sql.DB, error) {
	db, err := database.Connect("mysql", envPrefix)
	if err != nil {
		return nil, err
	}

	if err := runSQL("cleanup_db"); err != nil {
		return nil, err
	}

	if err := runSQL("setup_db"); err != nil {
		return nil, err
	}

	return db, nil
}

func teardown(db *sql.DB) error {
	if err := runSQL("cleanup_db"); err != nil {
		return err
	}

	db.Close()

	return nil
}

func runSQL(name string) error {
	host := os.Getenv(envPrefix + "DB_HOST")
	user := os.Getenv(envPrefix + "DB_USER")
	dbName := os.Getenv(envPrefix + "DB_NAME")
	password := os.Getenv(envPrefix + "DB_PASS")

	absPath, err := filepath.Abs("./testdata/" + name + ".sql")
	if err != nil {
		return err
	}

	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return err
	}

	cmd := exec.Command("mysql", "-h", host, "-u", user, dbName, "--password="+password, "-e", "source "+absPath)

	return cmd.Run()
}
