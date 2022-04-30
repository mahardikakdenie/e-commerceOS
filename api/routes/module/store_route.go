package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func StoreRoute(v1 *gin.RouterGroup, c *controller.StoreController, middleware gin.HandlerFunc) {
	store := v1.Group("store").Use(middleware)
	store.GET("/", c.FindAll)
	store.POST("/", c.Created)
	store.GET("/:id", c.Show)
	store.PATCH("/:id", c.Update)
	store.DELETE("/:id", c.Delete)
}
