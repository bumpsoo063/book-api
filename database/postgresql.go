package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var PQ *sql.DB = nil

func ConnectPq() (*sql.DB, error) {
	if PQ == nil {
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		connStr := fmt.Sprintf("host=%s dbname=%s user=%s port=%d password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), port, os.Getenv("DB_PASSWORD"))
		PQ, err = sql.Open("postgres", connStr)
		defer PQ.Close()
	}
	return PQ, nil
}
