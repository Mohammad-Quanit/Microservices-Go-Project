package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mohammad-quanit/Go-Microservices-App/handlers"
)

func main() {
	// A simple http1.1 web server
	port := 8000

	http.HandleFunc("/hello", handlers.HelloHttp)

	log.Printf("Server running on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
