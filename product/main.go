package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Create a new gin instance
	r := gin.Default()
	r.Use(gin.Logger())
	l := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddr := os.Getenv("PRODUCT_API_PORT")

	// Create an HTTP server instance
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		ErrorLog:     l,                 // set the logger for the server
	}

	go func() {
		l.Println("Server Started")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			os.Exit(1)
		}

	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 10 seconds for current operations to complete
	tCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure cancel function is called to avoid context leak

	if err := srv.Shutdown(tCtx); err != nil {
		log.Fatal("Server Shutdown Problem:", err)
	}
}

// productHandler := handlers.NewProducts(l)

// // Create goilla mux router instance
// r := mux.NewRouter()

// getRouter := r.Methods(http.MethodGet).Subrouter()
// getRouter.HandleFunc("/products", productHandler.GetAll)

// postRouter := r.Methods(http.MethodPost).Subrouter()
// postRouter.HandleFunc("/product", productHandler.Create)
// postRouter.Use(productHandler.ValidationMiddleware)

// putRouter := r.Methods(http.MethodPut).Subrouter()
// putRouter.HandleFunc("/product/{id:[0-9]+}", productHandler.Update)
// putRouter.Use(productHandler.ValidationMiddleware)

// s := http.Server{
// 	Addr:         ":9090",           // configure the bind address
// 	Handler:      r,                 // set the default handler
// 	IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
// 	ReadTimeout:  1 * time.Second,   // max time to read request from the client
// 	WriteTimeout: 10 * time.Second,  // max time to write response to the client
// 	ErrorLog:     l,                 // set the logger for the server
// }