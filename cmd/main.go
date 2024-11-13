package main

import (
	"log"

	server "github.com/DanielDDHM/Hire-Go/config"
	"github.com/DanielDDHM/Hire-Go/internal/database"
)

func main() {
	server := server.App()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	server.Run(db)
}
