package server

import (
	"log"
	"net/http"
	"os"

	"github.com/201-tech/Hire-Go/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	port   string
	server *gin.Engine
}

func App() Server {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error When Reload .env file")
	}

	PORT := os.Getenv("APP_PORT")
	return Server{
		port:   PORT,
		server: gin.Default(),
	}
}

func (s *Server) Run(db *gorm.DB) {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "I'm Alive",
			})
		})
		handler.User(v1.Group("/user"), db)
		handler.Role(v1.Group("/role"), db)
	}
	log.Print("Server is on port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
