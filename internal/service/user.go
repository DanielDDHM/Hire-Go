package service

import (
	"net/http"

	models "github.com/201-tech/Hire-Go/internal/dto"
	utils "github.com/201-tech/Hire-Go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": validationErr.Error()})
			return
		}

		hashedPassword, err := utils.HashPassword(user.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to hash password"})
			return
		}

		user.Password = hashedPassword

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user, "status": "success"})
	}
}

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User

		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
			return
		}

		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": users, "status": "success"})
	}
}
