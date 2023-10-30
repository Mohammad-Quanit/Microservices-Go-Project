package models

import (
	"gorm.io/gorm"
)

// swagger:model Product
type Product struct {
	gorm.Model
	Name           string  `json:"name" binding:"required"`
	Description    string  `json:"description"`
	Price          float32 `json:"price" binding:"required,gte=0"`
	Category       string  `json:"category" binding:"required"`
	Stock_Quantity float32 `json:"stock_quantity" binding:"required"`
	CreatedOn      string  `json:"-"`
	UpdatedOn      string  `json:"-"`
	DeletedOn      string  `json:"-"`
}

type Products []*Product
