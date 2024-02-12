package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Connect(driverName, prefix string) (*sql.DB, error) {
	switch strings.ToLower(driverName) {
	case "mysql":
		return connectToMySQL(prefix)
	case "postgres":
		return connectToPostgreSQL(prefix)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driverName)
	}
}

func connectToMySQL(prefix string) (*sql.DB, error) {
	loc, err := time.LoadLocation(os.Getenv(prefix + "DB_TIMEZONE"))
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv(prefix+"DB_HOST"), os.Getenv(prefix+"DB_PORT"))
	allowNativePasswords := strings.ToLower(os.Getenv(prefix + "DB_ALLOW_NATIVE_PASSWORDS"))

	config := mysql.Config{
		DBName:               os.Getenv(prefix + "DB_NAME"),
		User:                 os.Getenv(prefix + "DB_USER"),
		Passwd:               os.Getenv(prefix + "DB_PASS"),
		Addr:                 addr,
		Net:                  os.Getenv(prefix + "DB_PROTOCOL"),
		ParseTime:            true,
		Collation:            os.Getenv(prefix + "DB_COLLATION"),
		Loc:                  loc,
		AllowNativePasswords: allowNativePasswords == "true",
	}

	return sql.Open("mysql", config.FormatDSN())
}

func connectToPostgreSQL(prefix string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv(prefix+"DB_HOST"),
		os.Getenv(prefix+"DB_PORT"),
		os.Getenv(prefix+"DB_USER"),
		os.Getenv(prefix+"DB_NAME"),
		os.Getenv(prefix+"DB_PASS"),
		os.Getenv(prefix+"DB_SSL_MODE"),
	)

	return sql.Open("postgres", dsn)
}
