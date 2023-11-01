package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammad-quanit/Go-Microservices-App/product/middlewares"
	"github.com/mohammad-quanit/Go-Microservices-App/product/models"
	"github.com/mohammad-quanit/Go-Microservices-App/product/routes"
	"github.com/mohammad-quanit/Go-Microservices-App/product/utils"
)

func main() {
	// Initialize Gin router and setup middleware
	router := gin.Default()
	router.Use(middlewares.DefaultStructuredLogger())
	router.Use(middlewares.IOLogger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")

	// Initialize a standard logger
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	// Extract server configuration from environment variables
	serverAddr := os.Getenv("PRODUCT_API_PORT")
	dbConfig := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Create an HTTP server instance
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     logger,
	}

	// Initialize the database
	models.InitDB(dbConfig)

	// Load the routes
	routes.ProductRoutes(v1)

	// Start the server in a goroutine
	go func() {
		if err := router.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start: ", err)
			os.Exit(1)
		}
		// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	logger.Fatalf("Server startup failed: %s", err)
		// 	os.Exit(1)
		// }
	}()

	// Graceful shutdown of the server
	utils.GracefulShutdown(srv)
}
