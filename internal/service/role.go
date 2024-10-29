package service

import (
	"net/http"

	models "github.com/201-tech/Hire-Go/internal/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role models.Role

		if err := c.Bind(&role); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error(), "location": "bind"})
			return
		}

		if validationErr := validate.Struct(&role); validationErr != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": validationErr.Error(), "location": "validation"})
			return
		}

		if err := db.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error(), "location": "db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": role, "status": "success"})
	}
}

func GetRoles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []models.Role

		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
			return
		}

		if err := db.Find(&roles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": roles, "status": "success"})
	}
}
