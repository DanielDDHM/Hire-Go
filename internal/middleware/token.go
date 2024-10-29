package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("minha_chave_secreta")

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}
