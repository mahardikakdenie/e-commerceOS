package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func TshirtRoute(v1 *gin.RouterGroup, c *controller.TShirtController, middleware gin.HandlerFunc) {
	tshirt := v1.Group("tshirt").Use(middleware)
	tshirt.GET("/", c.FindAll)
	tshirt.POST("/", c.Store)
	tshirt.GET("/:id", c.Show)
	tshirt.PATCH("/:id", c.Updated)
	tshirt.DELETE("/:id", c.Destroy)
}
