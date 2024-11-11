package handler

import (
	"github.com/DanielDDHM/Hire-Go/internal/middleware"
	"github.com/DanielDDHM/Hire-Go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(routes *gin.RouterGroup, db *gorm.DB) {

	routes.POST("/", service.CreateUser(db))

	protectedRoutes := routes.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware())
	protectedRoutes.GET("/", service.GetUsers(db))
}
