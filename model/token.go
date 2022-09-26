package model

import (
	"github.com/google/uuid"
)

type Token struct {
	AccessToken   string
	RefreshToken  string `json:"refresh_token"`
	AccessUuid    uuid.UUID
	RefreshUuid   uuid.UUID
	AccessExpire  int64
	RefreshExpire int64
}
