package service

import (
	"net/http"

	"github.com/DanielDDHM/Hire-Go/internal/models"
	"github.com/DanielDDHM/Hire-Go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAsBooker(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var booker models.Booker
		var role models.Role

		if err := db.First(&role, "name = ?", "Booker").Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if !utils.BindAndValidate(c, &booker) {
			return
		}

		if err := db.Model(&models.Booker{}).
			Create(map[string]interface{}{
				"user_id": booker.UserId,
				"name":    booker.Name,
				"address": booker.Address,
				"city":    booker.City,
				"country": booker.Country,
				"phone":   booker.Phone,
			}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&models.User{}).
			Where("id = ?", booker.UserId).
			Update("role_id", role.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": booker, "status": "success"})
	}
}
