package server

import (
	"log"
	"net/http"
	"os"

	"github.com/DanielDDHM/Hire-Go/internal/handler"
	"github.com/DanielDDHM/Hire-Go/internal/middleware"
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
		roleGroup := v1.Group("/role")
		roleGroup.Use(middleware.AuthMiddleware(db, "Admin"))
		handler.Role(roleGroup, db)

		handler.User(v1.Group("/user"), db)
		handler.Auth(v1.Group("/auth"), db)

		v1.GET("/ws", func(ctx *gin.Context) {
			HandleWebSocket(ctx.Writer, ctx.Request)
		})
	}
	log.Print("Server is on port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
