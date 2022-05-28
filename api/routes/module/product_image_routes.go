package module

import (
	controller "api/modules/product/image/controller"

	"github.com/gin-gonic/gin"
)

func ProductImageRoute(v1 *gin.RouterGroup, c *controller.ImageController, middleware ...gin.HandlerFunc) {
	productImage := v1.Group("product-image").Use(middleware[0])
	productImage.POST("", middleware[1], c.Store)
	productImage.GET(":product_id", c.Index)
}
