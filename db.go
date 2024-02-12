package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	jst, err := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	allowNativePasswords := strings.ToLower(os.Getenv("DB_ALLOW_NATIVE_PASSWORDS"))

	config := mysql.Config{
		DBName:               os.Getenv("DB_NAME"),
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Addr:                 addr,
		Net:                  os.Getenv("DB_PROTOCOL"),
		ParseTime:            true,
		Collation:            os.Getenv("DB_COLLATION"),
		Loc:                  jst,
		AllowNativePasswords: allowNativePasswords == "true",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
