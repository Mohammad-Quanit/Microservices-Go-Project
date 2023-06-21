// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammad-quanit/Go-Microservices-App/data"
)

type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET requests")

	// fetch the products from the dummy datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to convert to JSON", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST requests")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod %#v", product)
	data.AddProduct(product)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get ID from request
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT requests", id)

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product Not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) ProductValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR]: Unable to unmarshall JSON", err)
			http.Error(w, "Unable to unmarshall JSON", http.StatusBadRequest)
			return
		}

		//validate the product for sanitization
		err = product.Validate()
		if err != nil {
			p.l.Println("[ERROR]: validating product", err)
			http.Error(w, fmt.Sprintf("error validating product %v", err), http.StatusBadRequest)
			return
		}

		// add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		// call the next handler, which can be another middleware int the chain, or the final handler
		next.ServeHTTP(w, r)
	})
}
