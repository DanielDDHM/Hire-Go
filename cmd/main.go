package main

import (
	server "github.com/201-tech/Hire-Go/config"
	"github.com/201-tech/Hire-Go/internal/database"
)

func main() {
	server := server.App()
	db := database.ConnectDB()
	server.Run(db)
}
