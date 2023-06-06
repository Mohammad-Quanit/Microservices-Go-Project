package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mohammad-quanit/Go-Microservices-App/handlers"
)

// "context"
// "log"
// "net/http"
// "os"
// "os/signal"
// "time"
func main() {

	// // http.DefaultServeMux.HandleFunc("/", hh.ServeHttp)
	// // http.HandleFunc("/", hh.ServeHttp)
	// // http.ListenAndServe(":9001", sm)

	// s := &http.Server{
	// 	Addr:         ":9001",
	// 	Handler:      sm,
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// }

	// go func() {
	// 	err := s.ListenAndServe()
	// 	if err != nil {
	// 		l.Fatal(err)
	// 	}
	// }()

	// sigChan := make(chan os.Signal)
	// signal.Notify(sigChan, os.Interrupt)
	// signal.Notify(sigChan, os.Kill)

	// sig := <-sigChan
	// l.Println("Recieved terminate, graceful shutdown", sig)

	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// s.Shutdown(ctx)

	l := log.New(os.Stdout, "Microservices-Go-Project", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.HandleFunc("/", hh.ServeHttp)

	// Listen for connections on all ip addresses (0.0.0.0) & port 9090
	log.Println("Server Starting...")
	if err := http.ListenAndServe(":9090", sm); err != nil {
		log.Fatal(err)
	}

}
