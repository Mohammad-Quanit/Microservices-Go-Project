package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/mohammad-quanit/auth/models"
)

// takes a JWT token as input and returns the claims contained in it. Claims are a set of key-value pairs that
// represent the information being transmitted between parties.

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("52E813EFDC167DA24970F9BDAA9A9B0948E687CA754A29B8615C78D3D80E2F84"), nil
	})

	if err != nil {
		fmt.Println("this is err...", err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	fmt.Println(ok)
	if !ok {
		return nil, err
	}

	return claims, nil
}
