package handler

import (
	service "github.com/201-tech/Hire-Go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(routes *gin.RouterGroup, db *gorm.DB) {
	routes.POST("/login", service.LoginUser(db))
}
