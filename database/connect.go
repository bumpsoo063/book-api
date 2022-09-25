package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s port=%d password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), port, os.Getenv("DB_PASSWORD"))
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
