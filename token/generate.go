package token

import (
	"os"
	"strconv"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// if f == true, only access token
// false -> both will be generdated
func Generate(userId string, f bool) (model.Token, error) {
	var token model.Token
	secret := os.Getenv("JWT_SECRET")
	minutes, err := strconv.Atoi(os.Getenv("JWT_SECRET_EXPIRE_MINUTES"))
	if err != nil {
		minutes = 10
	}
	id, err := uuid.NewUUID()
	token.AccessUuid = id.String()
	if err != nil {
		return token, err
	}
	token.AccessExpire = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["access_uuid"] = token.AccessUuid
	claims["user_id"] = userId
	claims["exp"] = token.AccessExpire
	token.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return token, err
	}
	if f == true {
		return token, nil
	}
	rsecret := os.Getenv("JWT_REFRESH_SECRET")
	days, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRE_DAYS"))
	if err != nil {
		days = 3
	}
	token.RefreshExpire = time.Now().Add(time.Hour * 24 * time.Duration(days)).Unix()
	if err != nil {
		return token, err
	}
	id, err = uuid.NewUUID()
	token.RefreshUuid = id.String()
	claims = jwt.MapClaims{}
	claims["authorized"] = true
	claims["refresh_uuid"] = token.RefreshUuid
	claims["user_id"] = userId
	claims["exp"] = token.RefreshExpire
	token.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(rsecret))
	if err != nil {
		return token, err
	}
	return token, nil
}
