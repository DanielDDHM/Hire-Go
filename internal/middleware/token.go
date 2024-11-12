package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/DanielDDHM/Hire-Go/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func getJwtKey() ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error when loading environment variables: %v", err)
	}

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return nil, fmt.Errorf("JWT_KEY not defined")
	}

	return []byte(jwtKey), nil
}

func GenerateJWT(user *models.User) (string, error) {
	jwtKey, err := getJwtKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]interface{}{
			"id":      user.Id,
			"name":    user.Name,
			"role_id": user.RoleId,
		},
		"exp": time.Now().Add(30 * time.Minute).Unix(),
	})
	return token.SignedString(jwtKey)
}

func DecodeToken(tokenString string) (*models.User, error) {
	jwtKey, err := getJwtKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error when decoding: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userMap, exists := claims["user"].(map[string]interface{}); exists {
			user := &models.User{
				Id:     int(userMap["id"].(float64)),
				Name:   userMap["name"].(string),
				RoleId: int(userMap["role_id"].(float64)),
			}
			return user, nil
		}
		return nil, fmt.Errorf("user not found in token")
	}
	return nil, fmt.Errorf("invalid token")
}
