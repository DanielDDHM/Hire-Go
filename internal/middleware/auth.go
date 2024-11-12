package middleware

import (
	"fmt"
	"net/http"

	"github.com/DanielDDHM/Hire-Go/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		user, err := DecodeToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		var userFind models.User

		if err := db.First(&userFind, "id = ?", user.Id).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		_, err = ValidateRole(db, user, role)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role invalid"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func ValidateRole(db *gorm.DB, user *models.User, role string) (bool, error) {

	if role == "all" {
		return true, nil
	}

	var roleFind *models.Role

	if err := db.First(&roleFind, "id = ?", user.RoleId).Error; err != nil {
		return false, err
	}

	if roleFind.Name != role {
		return false, fmt.Errorf("user doesn't have authorization")
	}

	return true, nil
}
