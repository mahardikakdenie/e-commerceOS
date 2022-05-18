package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func CustomerAuthRoute(v1 *gin.RouterGroup, c *controller.CustomerAuthController, middleware gin.HandlerFunc, cors gin.HandlerFunc) {
	customerAuth := v1.Group("customer_auth")
	customerAuth.GET("/", middleware, c.Index)
	customerAuth.POST("/login", cors, c.Login)
	customerAuth.POST("/register", cors, c.Register)
	customerAuth.POST("/logout", cors, middleware, c.Logout)
}
