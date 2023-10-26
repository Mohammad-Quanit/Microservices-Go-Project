package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/auth/utils"
)

func Home(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}
