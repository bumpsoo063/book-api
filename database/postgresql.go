package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var pq *sql.DB

func ConnectPq() error {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s port=%d password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), port, os.Getenv("DB_PASSWORD"))
	pq, err = sql.Open("postgres", connStr)
	return nil
}

func GetPq() *sql.DB {
	return pq
}
