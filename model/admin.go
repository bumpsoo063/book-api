package model

import (
	"github.com/google/uuid"
)

type Admin struct {
	Id       uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
