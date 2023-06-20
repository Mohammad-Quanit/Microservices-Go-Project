package data

import "testing"

func TestProduct_Validate(t *testing.T) {
	p := &Product{
		Name:  "Quanit",
		Price: 10,
		SKU:   "abc-www",
	}
	err := p.Validate()

	if err != nil {
		t.Errorf("Products.Validate() = %v", err)
	}
}
