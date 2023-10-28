package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/Go-Microservices-App/product/controllers"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.AddProduct)
	r.GET("/products/:id", controllers.GetProduct)
	r.DELETE("/products/:id", controllers.RemoveProduct)
}
