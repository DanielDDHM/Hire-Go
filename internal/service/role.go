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

		// Validate the request body
		if err := c.Bind(&role); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error(), "location": "bind"})
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate.Struct(&role); validationErr != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": validationErr.Error(), "location": "validation"})
			return
		}

		// Save the user to the database
		if err := db.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error(), "location": "db"})
			return
		}

		c.JSON(http.StatusOK, role)
	}
}

func GetRoles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []models.Role

		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
			return
		}

		// Fetch all users from the database
		if err := db.Find(&roles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, roles)
	}
}
