package gateway

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/Go-Microservices-App/product/middlewares"
)

func main() {
	// Initialize your API gateway (e.g., using the Gin framework)
	// Initialize Gin router and setup middleware
	router := gin.Default()
	router.Use(middlewares.DefaultStructuredLogger())
	router.Use(middlewares.IOLogger())
	router.Use(gin.Recovery())

	// Define your routes and their handlers
	// router.Handle("/v1/", AuthMiddleware()) // Routes that need authentication

	// Start your server
	if err := router.Run(); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	log.Fatal("Server failed to start: ", err)
	// }
}

// Middleware function to validate JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(c.Writer, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Extract JWT token from the Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate JWT (sample validation, replace with your actual JWT validation logic)
		if !isValidToken(token) {
			http.Error(c.Writer, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, proceed to the next handler (product handler, in this case)
		c.Next()
		// next.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler for Product microservice
// func ProductHandler(w http.ResponseWriter, r *http.Request) {
// 	// Handle requests specific to the Product microservice
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Product service response"))
// }

// Validate JWT (replace with your actual validation logic)
func isValidToken(token string) bool {
	// Your JWT validation logic (verify signature, expiration, claims, etc.)
	// Return true if the token is valid, false otherwise
	return true
}
