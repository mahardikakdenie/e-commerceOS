package module

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func RoleRoute(v1 *gin.RouterGroup, c *controller.RoleController, middleware gin.HandlerFunc) {
	role := v1.Group("role").Use(middleware)
	role.GET("/", c.FindAll)
}
