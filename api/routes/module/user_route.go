package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(v1 *gin.RouterGroup, c *controller.UserController, middleware gin.HandlerFunc) {
	user := v1.Group("user").Use(middleware)
	user.GET("/", c.FindAll)
	user.GET("/me", c.Me)
} // UserRoute
