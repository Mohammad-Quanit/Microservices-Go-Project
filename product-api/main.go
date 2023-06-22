package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/mohammad-quanit/Go-Microservices-App/handlers"
)

func main() {
	l := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	productHandler := handlers.NewProducts(l)

	// Create goilla mux router instance
	r := mux.NewRouter()

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetAll)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.Create)
	postRouter.Use(productHandler.ValidationMiddleware)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Update)
	putRouter.Use(productHandler.ValidationMiddleware)

	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      r,                 // set the default handler
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		ErrorLog:     l,                 // set the logger for the server
	}

	go func() {
		// Listen for connections on all ip addresses (0.0.0.0) & port 9090
		l.Println("Server Starting...")
		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure cancel function is called to avoid context leak
	s.Shutdown(tCtx)
}
