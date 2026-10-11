package main

import (
	"log"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inisialisasi database
	database.InitDB()

	// Jalankan router
	router := routes.SetupRouter()
	router.Run(":3000")
}
