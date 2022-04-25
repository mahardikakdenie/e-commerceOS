package controller

import (
	"api/auth"
	"api/helper"
	"api/user"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authRepository auth.Service
}

func NewAuthController(authRepository auth.Service) *AuthController {
	return &AuthController{authRepository}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	var token_request auth.AuthRequest

	user, err := controller.authRepository.FindByEmail(email)

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

	check := helper.CheckPasswordHash(password, user.Password)
	if !check {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": "Password is incorrect",
			},
			"data": gin.H{},
		})
		return
	}

	randomToken := make([]byte, 32)
	_, errRand := rand.Read(randomToken)

	if errRand != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{},
		})
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)

	token_request = auth.AuthRequest{
		UserId:    int(user.ID),
		AuthToken: authToken,
	}

	auth, err := controller.authRepository.Login(token_request)

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{},
		"data": auth.AuthToken,
	})
}

func (controller *AuthController) Register(ctx *gin.Context) {
	var userRequest user.UserRequest

	password := ctx.PostForm("password")

	hash_password, _ := helper.HashPassword(password)
	roleId := ctx.PostForm("role_id")
	roleIdString, _ := strconv.Atoi(roleId)

	userRequest = user.UserRequest{
		Name:     ctx.PostForm("name"),
		Email:    ctx.PostForm("email"),
		Password: hash_password,
		RoleId:   roleIdString,
	}

	user, err := controller.authRepository.Register(userRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": "User already exists",
				"errors ": err.Error(),
			},
			"data": gin.H{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "User created",
		},
		"data": user,
	})
}