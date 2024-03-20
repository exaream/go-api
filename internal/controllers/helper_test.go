package controllers_test

import (
	"database/sql"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/exaream/go-api/internal/config"
	"github.com/exaream/go-api/internal/database"
)

const (
	dotenvPath = "../../.env"
	sqlDir     = "../../_develop/mysql/sql"
)

func Setup() (*config.Config, *sql.DB, error) {
	config.Load(dotenvPath)
	cfg, err := config.Get()
	if err != nil {
		return nil, nil, err
	}

	db, err := database.Connect("mysql", &cfg.TestDB)
	if err != nil {
		return nil, nil, err
	}

	if err := RunSQL(&cfg.TestDB, sqlDir+"/create_tables.sql"); err != nil {
		return nil, nil, err
	}

	if err := RunSQL(&cfg.TestDB, sqlDir+"/insert_into_tables.sql"); err != nil {
		return nil, nil, err
	}

	return cfg, db, nil
}

func Teardown(cfg *config.Config) error {
	return RunSQL(&cfg.TestDB, sqlDir+"/clear_tables.sql")
}

func RunSQL(cfg *config.DB, filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return err
	}

	cmd := exec.Command("mysql",
		"-h", cfg.Host,
		"-u", cfg.User,
		"-P", cfg.Port,
		cfg.Name,
		"--password="+cfg.Password,
		"-e", "source "+absPath)

	return cmd.Run()
}
