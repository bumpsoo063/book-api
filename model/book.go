package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID `uri:"uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
}
