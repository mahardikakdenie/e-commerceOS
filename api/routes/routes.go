package routes

import (
	"api/auth"
	"api/controller"
	"api/middleware"
	"api/routes/module"
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

	MyMiddleware := middleware.MyMiddleware(authService)

	v1 := router.Group("v1")

	// routes
	module.AuthRoute(v1, authController)
	module.UserRoute(v1, userController, MyMiddleware)

}
