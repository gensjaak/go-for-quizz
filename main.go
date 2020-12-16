package main

import (
	"gensjaak/go-for-quizz/src/models"
	"gensjaak/go-for-quizz/src/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()

	// Log an error if error encountered
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read specific variables
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USERNAME")
	pass := os.Getenv("DATABASE_PASSWORD")
	apiPort := os.Getenv("API_PORT")

	// Attempt to connect to database - 5 tries to fallback and exit process
	models.Connect(host, port, dbName, user, pass, 5)

	// After router configuration, run the app
	router.Configure().Run(":" + apiPort)
}
