package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/Go-Microservices-App/product/models"
)

func GetProducts(c *gin.Context) {
	var products models.Products
	models.DB.Find(&products)
	c.JSON(200, gin.H{"data": products})
}
