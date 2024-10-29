package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/201-tech/Hire-Go/internal/utils"
)

type PasswordRequest struct {
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

func ValidatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PasswordRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Senha e hash são obrigatórios"})
			c.Abort()
			return
		}

		valid := utils.CheckPasswordHash(req.Password, req.HashedPassword)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha incorreta"})
			c.Abort()
			return
		}

		c.Next()
	}
}
