package token

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateToken() (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	minutes, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES"))
	if err != nil {
		minutes = 15
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ret, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return ret, nil
}
