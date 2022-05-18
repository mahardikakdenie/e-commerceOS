package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func CartRoute(v1 *gin.RouterGroup, c *controller.CartController, middleware gin.HandlerFunc) {
	cart := v1.Group("cart").Use(middleware)
	cart.GET("", c.Index)
	cart.POST("", c.Store)
	cart.GET("/:id", c.Show)
	cart.PATCH("/:id", c.Update)
	cart.DELETE("/:id", c.Destroy)
}
