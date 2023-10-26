package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammad-quanit/auth/models"
	"github.com/mohammad-quanit/auth/routes"
)

func main() {
	// Create a new gin instance
	r := gin.Default()
	r.Use(gin.Logger())

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddr := os.Getenv("PORT")

	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// initialize DB
	models.InitDB(config)

	// Load the routes
	routes.AuthRoutes(r)

	// Run the server
	r.Run(serverAddr)
}
