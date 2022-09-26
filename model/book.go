package model

import (
	"time"
)

type Book struct {
	Id        string `uri:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
}
