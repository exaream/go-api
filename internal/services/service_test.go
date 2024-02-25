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

	"github.com/exaream/go-api/internal/database"
	"github.com/exaream/go-api/internal/services"
)

const (
	envPrefix   = "TEST_"
	sqlDir      = "../../_develop/mysql/sql"
	testdataDir = "./testdata"
)

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

	if err := runSQL(sqlDir + "/drop_tables.sql"); err != nil {
		return nil, err
	}

	if err := runSQL(sqlDir + "/create_tables.sql"); err != nil {
		return nil, err
	}

	if err := runSQL(testdataDir + "/insert_into_tables.sql"); err != nil {
		return nil, err
	}

	return db, nil
}

func teardown(db *sql.DB) error {
	if err := runSQL(sqlDir + "/drop_tables.sql"); err != nil {
		return err
	}

	db.Close()

	return nil
}

func runSQL(filename string) error {
	filename = filepath.Clean(filename)
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	host := os.Getenv(envPrefix + "DB_HOST")
	user := os.Getenv(envPrefix + "DB_USER")
	dbName := os.Getenv(envPrefix + "DB_NAME")
	password := os.Getenv(envPrefix + "DB_PASS")

	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return err
	}

	cmd := exec.Command("mysql", "-h", host, "-u", user, dbName, "--password="+password, "-e", "source "+absPath)

	return cmd.Run()
}
