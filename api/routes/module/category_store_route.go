package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func CategoryStoreRoute(v1 *gin.RouterGroup, c *controller.CategoryStoreController, middleware gin.HandlerFunc, cors gin.HandlerFunc) {
	category_store := v1.Group("category-store").Use(middleware)
	category_store.GET("/", cors, c.Index)
	category_store.POST("/", c.Created)
	category_store.GET("/:id", c.Show)
	category_store.PATCH("/:id", c.Updated)
	category_store.DELETE("/:id", c.Deleted)
}
