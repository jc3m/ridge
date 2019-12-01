package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	return sql.Open("mysql", dsn)
}
