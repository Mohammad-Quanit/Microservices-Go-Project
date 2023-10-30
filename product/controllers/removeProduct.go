package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/Go-Microservices-App/product/models"
)

// DELETE /products/:id
// Delete a product
func RemoveProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
