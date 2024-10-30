package handler

import (
	"github.com/201-tech/Hire-Go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Role(routes *gin.RouterGroup, db *gorm.DB) {
	routes.POST("/", service.CreateRole(db))
	routes.GET("/", service.GetRoles(db))
}
