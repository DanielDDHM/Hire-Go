package handler

import (
	"github.com/201-tech/Hire-Go/internal/middleware"
	service "github.com/201-tech/Hire-Go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(routes *gin.RouterGroup, db *gorm.DB) {

	routes.POST("/", service.CreateUser(db))

	protectedRoutes := routes.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware())
	protectedRoutes.GET("/", service.GetUsers(db))
}
