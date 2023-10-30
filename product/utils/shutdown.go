package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Trap sigterm or interupt and gracefully shutdown the server
// waiting max 10 seconds for current operations to complete
func GracefulShutdown(srv *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Recieved terminate, graceful shutdown", sig)

	tCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure cancel function is called to avoid context leak

	if err := srv.Shutdown(tCtx); err != nil {
		log.Fatal("Server Shutdown Problem:", err)
	}
}
