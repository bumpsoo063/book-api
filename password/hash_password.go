package password

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
)

func HashPassword(str string) string {
	salt := []byte(os.Getenv("SHA_SALT"))
	password := []byte(str)
	hasher := sha512.New()
	password = append(password, salt...)
	hasher.Write(password)
	hashedPasswordBytes := hasher.Sum(nil)
	str = hex.EncodeToString(hashedPasswordBytes)
	return str
}
