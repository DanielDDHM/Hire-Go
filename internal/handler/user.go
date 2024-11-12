package handler

import (
	"github.com/DanielDDHM/Hire-Go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(routes *gin.RouterGroup, db *gorm.DB) {

	routes.POST("/", service.CreateUser(db))
	routes.GET("/", service.GetUsers(db))
}
