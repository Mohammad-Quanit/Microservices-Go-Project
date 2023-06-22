package handlers

import (
	"net/http"

	"github.com/mohammad-quanit/Go-Microservices-App/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//  200: productsResponse

// GetAll handles GET requests and returns all current products
func (p *Products) GetAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	// fetch the products from the dummy datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to convert to JSON", http.StatusInternalServerError)
		return
	}
}
