package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/auth/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/home", controllers.Home)
	r.GET("/premium", controllers.Premium)
	r.GET("/logout", controllers.Logout)
}
