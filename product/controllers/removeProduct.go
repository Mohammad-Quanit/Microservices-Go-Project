package controllers

import "github.com/gin-gonic/gin"

func RemoveProduct(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "pong"})
}
