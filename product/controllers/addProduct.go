package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/Go-Microservices-App/product/models"
)

func AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// product = models.Product{Name: product.Name, Description: product.Description, Price: product.Price, SKU: product.SKU}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
