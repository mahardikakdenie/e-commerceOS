package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func CategoryCustomerRoute(v1 *gin.RouterGroup, c *controller.CategoryCustomerController, cors gin.HandlerFunc) {
	category_customer := v1.Group("category-customer").Use(cors)
	category_customer.GET("", c.Index)
}
