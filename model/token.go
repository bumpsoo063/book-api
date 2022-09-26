package model

type Token struct {
	AccessToken   string
	RefreshToken  string `json:"refresh_token"`
	AccessUuid    string
	RefreshUuid   string
	AccessExpire  int64
	RefreshExpire int64
}
