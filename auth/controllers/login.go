package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/auth/models"
	"github.com/mohammad-quanit/auth/utils"
)

func Login(c *gin.Context) {

	// Recieve User Information
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if user/email doesn't exist
	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user doesn't exist"})
		return
	}

	// Check if user password is invalid
	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	// Create Claim with exp timne and create jwt token
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	// Set token string in Cookies
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in"})
}
