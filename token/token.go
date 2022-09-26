package token

import (
	"os"
	"strconv"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateToken(user uuid.UUID) (model.Token, error) {
	var token model.Token
	secret := os.Getenv("JWT_SECRET_KEY")
	minutes, err := strconv.Atoi(os.Getenv("JWT_SECRET_EXPIRE_MINUTES"))
	if err != nil {
		minutes = 10
	}
	token.AccessUuid, err = uuid.NewUUID()
	if err != nil {
		return token, err
	}
	token.AccessExpire = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["access_uuid"] = token.AccessUuid
	claims["user_id"] = user
	claims["exp"] = token.AccessExpire
	token.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(secret))
	if err != nil {
		return token, err
	}
	rsecret := os.Getenv("JWT_REFRESH_SECRET")
	days, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRE_DAYS"))
	if err != nil {
		days = 3
	}
	token.RefreshExpire = time.Now().Add(time.Hour * 24 * time.Duration(days)).Unix()
	token.RefreshUuid, err = uuid.NewUUID()
	if err != nil {
		return token, err
	}
	claims = jwt.MapClaims{}
	claims["authorized"] = true
	claims["refresh_uuid"] = token.RefreshUuid
	claims["user_id"] = user
	claims["exp"] = token.RefreshExpire
	token.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(rsecret))
	
	return token, nil
}
