package handlers

import (
	"net/http"

	"github.com/mohammad-quanit/Go-Microservices-App/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//
//	200: productsResponse
//
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

// // ListSingle handles GET requests
// func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
// 	id := getProductID(r)

// 	p.l.Println("[DEBUG] get record id", id)

// 	prod, err := data.GetProductByID(id)

// 	l := data.GetProducts()

// 	switch err {
// 	case nil:

// 	case data.ErrProductNotFound:
// 		p.l.Println("[ERROR] fetching product", err)

// 		rw.WriteHeader(http.StatusNotFound)
// 		l.ToJSON()
// 		return
// 	default:
// 		p.l.Println("[ERROR] fetching product", err)

// 		rw.WriteHeader(http.StatusInternalServerError)
// 		l.ToJSON(&GenericError{Message: err.Error()}, rw)
// 		return
// 	}

// 	err = l.ToJSON(prod, rw)
// 	if err != nil {
// 		// we should never be here but log the error just incase
// 		p.l.Println("[ERROR] serializing product", err)
// 	}
// }
