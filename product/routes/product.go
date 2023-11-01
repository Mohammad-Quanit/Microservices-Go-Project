package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/Go-Microservices-App/product/controllers"
)

func ProductRoutes(v *gin.RouterGroup) {
	v.GET("/")
	v.GET("/ping", controllers.Ping)
	v.GET("/products", controllers.GetProducts)
	v.POST("/products", controllers.AddProduct)
	v.GET("/products/:id", controllers.GetProduct)
	v.DELETE("/products/:id", controllers.RemoveProduct)
}
