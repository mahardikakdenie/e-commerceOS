package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func ProductStoreRoute(v1 *gin.RouterGroup, c *controller.ProductStoreController, middleware gin.HandlerFunc) {
	productStore := v1.Group("product-store").Use(middleware)
	productStore.GET("/", c.Index)
	productStore.GET("/:id", c.Show)
	productStore.POST("/", c.Created)
	productStore.PATCH("/:id", c.Updated)
	productStore.DELETE("/:id", c.Deleted)
}
