package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func ProductCustomerRoute(v1 *gin.RouterGroup, c *controller.ProductCustomerController, cors gin.HandlerFunc) {
	productCustomer := v1.Group("product-customer").Use(cors)
	productCustomer.GET("", c.Index)
	productCustomer.GET("/:slug", c.ProductByCategory)
	productCustomer.GET("/show/:id", c.Show)
}
