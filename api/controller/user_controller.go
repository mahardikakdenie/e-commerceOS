package controller

import (
	"api/middleware"
	"api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service user.Service
}

func NewController(service user.Service) *UserController {
	return &UserController{service}
}

func (controller *UserController) FindAll(c *gin.Context) {
	users, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"meta": gin.H{},
	})
}

func (controler *UserController) Me(ctx *gin.Context) {
	user, err := controler.service.Me(middleware.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": "User not found",
			},
			"data": gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
	})
}
