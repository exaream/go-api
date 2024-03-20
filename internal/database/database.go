package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/exaream/go-api/internal/config"
)

func Connect(driverName string, cfg *config.DB) (*sql.DB, error) {
	switch strings.ToLower(driverName) {
	case "mysql":
		return connectToMySQL(cfg)
	case "postgres":
		return connectToPostgreSQL(cfg)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driverName)
	}
}

func connectToMySQL(cfg *config.DB) (*sql.DB, error) {
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}

	c := &mysql.Config{
		Addr:                 fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		DBName:               cfg.Name,
		User:                 cfg.User,
		Passwd:               cfg.Password,
		Net:                  cfg.Protocol,
		ParseTime:            true,
		Collation:            cfg.Collation,
		Loc:                  loc,
		AllowNativePasswords: cfg.AllowNativePasswords,
	}

	return sql.Open("mysql", c.FormatDSN())
}

func connectToPostgreSQL(cfg *config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password, cfg.SSLMode,
	)

	return sql.Open("postgres", dsn)
}
