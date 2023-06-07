package handlers

import (
	"log"
	"net/http"

	"github.com/mohammad-quanit/Go-Microservices-App/data"
)

type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler
// interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	//catch all
	// if no method is satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

	// fetch the products from the datastore
	lp := data.GetProducts()
	// b, err := json.Marshal(lp)

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to convert to JSON", http.StatusInternalServerError)
		return
	}
	// w.Write(b)
}
