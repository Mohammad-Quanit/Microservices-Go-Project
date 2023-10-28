package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateSKU(f validator.FieldLevel) bool {
	// SKU is of format abc-def-ghi
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(f.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
