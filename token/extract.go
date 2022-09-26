package token

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func ExtractAccess(auth string) (string, error) {
	temp := strings.Split(auth, " ")
	var tokenString string
	if len(temp) == 2 {
		tokenString = temp[1]
	}
	secret := os.Getenv("JWT_SECRET")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		access := claims["access_uuid"].(string)
		return access, nil
	} else {
		return "", fmt.Errorf("token is not valid")
	}
}
