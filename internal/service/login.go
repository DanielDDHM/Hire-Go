package service

import (
	"net/http"

	"github.com/DanielDDHM/Hire-Go/internal/middleware"
	"github.com/DanielDDHM/Hire-Go/internal/models"
	"github.com/DanielDDHM/Hire-Go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Login

		if !utils.BindAndValidate(c, &req) {
			return
		}

		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "location": "db"})
			return
		}

		_, err := utils.CheckPasswordHash(req.Password, user.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "location": "CheckPasswordHash"})
			return
		}

		token, err := middleware.GenerateJWT(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "location": "GenerateJWT"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
