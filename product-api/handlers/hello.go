package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
func (h *Hello) ServeHttp(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Running Hello Handler")

	//read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body!", err)
		http.Error(w, "Unable to read request body!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", b)
}
