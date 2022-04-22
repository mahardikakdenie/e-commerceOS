package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoute(v1 *gin.RouterGroup, c *controller.OrderController, middleware gin.HandlerFunc) {
	order := v1.Group("order").Use(middleware)
	order.GET("/", c.Index)
	order.POST("/", c.Store)
	order.GET("/:id", c.Show)
	order.PATCH("/:id", c.Update)
	order.DELETE("/:id", c.Destroy)
}
