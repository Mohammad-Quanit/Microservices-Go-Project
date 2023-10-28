package controllers

import (
	"net/http"

	"github.com/mohammad-quanit/Go-Microservices-App/product/data"
)

// Create handles POST requests to add new products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST requests")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", product)
	data.AddProduct(product)
}
