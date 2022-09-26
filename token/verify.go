package token

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Parse(auth string) (*jwt.Token, error) {
	temp := strings.Split(auth, " ")
	var tokenString string
	if len(temp) == 2 {
		tokenString = temp[1]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractAccess(token *jwt.Token) (uuid.UUID, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		access := claims["access_uuid"].(uuid.UUID)
		return access, nil
	}
	return uuid.UUID{}, fmt.Errorf("token is not valid")
}