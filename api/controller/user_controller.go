package controller

import (
	"api/helper"
	"api/middleware"
	"api/modules/user"
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
	data, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userResponses []user.UserResponses
	for _, v := range data {
		userResponses = append(userResponses, user.UserResponses{
			ID:    uint(v.ID),
			Name:  v.Name,
			Email: v.Email,
			Auth:  v.Auth,
			Role:  v.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponses,
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

	userResponses := helper.UserResponses(user)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userResponses,
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
	})
}
