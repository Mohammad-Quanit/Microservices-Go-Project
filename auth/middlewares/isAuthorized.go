package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/auth/utils"
)

// The function IsAuthorized returns a Gin handler function,
// which is used to process HTTP requests and responses in a Gin web application.
func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")

		if err != nil {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
