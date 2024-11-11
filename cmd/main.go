package main

import (
	server "github.com/DanielDDHM/Hire-Go/config"
	"github.com/DanielDDHM/Hire-Go/internal/database"
)

func main() {
	server := server.App()
	db := database.ConnectDB()
	server.Run(db)
}
