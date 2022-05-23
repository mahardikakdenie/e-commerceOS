package module

import (
	cart_customer_controller "api/modules/cart/customer/controller"

	"github.com/gin-gonic/gin"
)

func CartCustomerRoute(v1 *gin.RouterGroup, controller *cart_customer_controller.CartsController, cors gin.HandlerFunc, middleware gin.HandlerFunc) {

	// controller := &cart_customer_controller.CartsController{
	// 	repo: cart_customer_service_interface.CartCussServiceInterface,
	// }

	cart_cuss_route := v1.Group("/cart_customer")

	cart_cuss_route.GET("/carts", cors, controller.Index)
	cart_cuss_route.POST("/carts/:store_id/:product_id", cors, middleware, controller.Store)
}
