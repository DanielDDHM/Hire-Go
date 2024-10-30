package service

import (
	"net/http"

	"github.com/201-tech/Hire-Go/internal/middleware"
	"github.com/201-tech/Hire-Go/internal/models"
	"github.com/201-tech/Hire-Go/internal/utils"
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

		if !utils.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password", "location": "CheckPasswordHash"})
			return
		}

		token, err := middleware.GenerateJWT(uint(user.Id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "location": "GenerateJWT"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
