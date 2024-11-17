package handler

import (
	"github.com/DanielDDHM/Hire-Go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Booker(routes *gin.RouterGroup, db *gorm.DB) {
	routes.POST("/", service.RegisterAsBooker(db))
}
