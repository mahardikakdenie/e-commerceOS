package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoute(v1 *gin.RouterGroup, c *controller.AuthController) {
	auth := v1.Group("auth")
	auth.POST("/login", c.Login)
	auth.POST("/register", c.Register)
}
