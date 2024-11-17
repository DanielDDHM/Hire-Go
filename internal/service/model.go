package service

import (
	"net/http"

	"github.com/DanielDDHM/Hire-Go/internal/models"
	"github.com/DanielDDHM/Hire-Go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func RegisterAsModel(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.Model
		var role models.Role

		if err := db.First(&role, "Name = ?", "Model").Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if !utils.BindAndValidate(c, &model) {
			return
		}

		if err := db.Model(&models.Model{}).
			Create(map[string]interface{}{
				"user_id": model.UserId,
				"name":    model.Name,
				"age":     model.Age,
				"height":  model.Height,
				"weight":  model.Weight,
				"bio":     model.Bio,
				"photos":  pq.Array(model.Photos),
				"address": model.Address,
				"city":    model.City,
				"country": model.Country,
			}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&models.User{}).
			Where("id = ?", model.UserId).
			Update("role_id", role.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": model, "status": "success"})
	}
}
