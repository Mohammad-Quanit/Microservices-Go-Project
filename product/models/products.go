package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/mohammad-quanit/Go-Microservices-App/product/utils"
	"gorm.io/gorm"
)

// swagger:model Product
type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
}

type Products []*Product

// Products Struct validation
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", utils.ValidateSKU)
	return validate.Struct(p)
}
