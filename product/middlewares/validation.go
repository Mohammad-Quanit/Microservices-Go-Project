package middlewares

func ValidationMiddleware() {}

// func ValidatePrice(f validator.FieldLevel) bool {
// 	// Price should be only positive numbers
// 	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
// 	matches := reg.FindAllString(f.Field().String(), -1)

// 	if len(matches) != 1 {
// 		return false
// 	}

// 	return true
// }
