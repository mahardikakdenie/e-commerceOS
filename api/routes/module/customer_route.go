package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func CustomerRoute(v1 *gin.RouterGroup, c *controller.CustomerController, middleware gin.HandlerFunc) {
	customer := v1.Group("customer").Use(middleware)
	customer.GET("/", c.Index)
	customer.POST("/", c.Store)
}
