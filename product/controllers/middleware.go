package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mohammad-quanit/Go-Microservices-App/product/data"
)

func (p *Products) ValidationMiddleware(next http.Handler) http.Handler {
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
