package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	return sql.Open("mysql", conn)
}
