package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammad-quanit/Go-Microservices-App/product/data"
)

func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
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
