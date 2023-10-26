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
	"log"

	"github.com/mohammad-quanit/Go-Microservices-App/data"
)

// ProductsResponseWrapper represents the response wrapper for products.
// swagger:response productsResponse
type ProductsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Products `json:"body"`
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
