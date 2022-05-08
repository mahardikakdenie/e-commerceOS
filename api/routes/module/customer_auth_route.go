package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func CustomerAuthRoute(v1 *gin.RouterGroup, c *controller.CustomerAuthController, middleware gin.HandlerFunc) {
	customerAuth := v1.Group("customer_auth")
	customerAuth.GET("/", middleware, c.Index)
	customerAuth.POST("/login", c.Login)
}
