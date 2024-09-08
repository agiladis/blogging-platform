package main

import (
	"blogging-platform/config"
	"blogging-platform/router"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// initiate DB
	DB, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := router.StartServer(DB.Debug()).Run(":3000"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
