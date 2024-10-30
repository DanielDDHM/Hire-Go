package service

import (
	"net/http"

	"github.com/201-tech/Hire-Go/internal/models"
	"github.com/201-tech/Hire-Go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role models.Role

		if !utils.BindAndValidate(c, &role) {
			return
		}

		if err := db.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "location": "db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": role, "status": "success"})
	}
}

func GetRoles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []models.Role

		if err := db.Find(&roles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "location": "db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": roles, "status": "success"})
	}
}
