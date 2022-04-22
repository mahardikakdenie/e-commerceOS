package routes

import (
	"api/auth"
	"api/controller"
	"api/middleware"
	orders "api/order"
	"api/role"
	"api/routes/module"
	"api/tshirt"

	"api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, router gin.IRouter) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := controller.NewController(userService)

	//

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authController := controller.NewAuthController(authService)

	//

	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)
	roleController := controller.NewRoleController(roleService)

	//

	tshirtRepository := tshirt.NewRepository(db)
	tshirtService := tshirt.NewService(tshirtRepository)
	tshirtController := controller.NewTShirtController(tshirtService)

	//

	orderRepository := orders.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	MyMiddleware := middleware.MyMiddleware(authService)

	v1 := router.Group("v1")

	// routes
	module.AuthRoute(v1, authController)
	module.UserRoute(v1, userController, MyMiddleware)
	module.RoleRoute(v1, roleController, MyMiddleware)
	module.TshirtRoute(v1, tshirtController, MyMiddleware)
	module.OrderRoute(v1, orderController, MyMiddleware)
}
