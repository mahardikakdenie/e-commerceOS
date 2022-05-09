package controller

import (
	"api/modules/role"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	repository role.Repository
}

func NewRoleController(repository role.Repository) *RoleController {
	return &RoleController{repository}
}

func (c *RoleController) FindAll(ctx *gin.Context) {
	roles, err := c.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": err.Error(),
			},
			"data": gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": roles,
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
	})
}
