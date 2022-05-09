package controller

import (
	"api/helper"
	"api/modules/customer_auth"
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type CustomerAuthController struct {
	service customer_auth.Service
}

func NewCustomerAuthController(service customer_auth.Service) *CustomerAuthController {
	return &CustomerAuthController{service}
}

func (c *CustomerAuthController) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	customer, err := c.service.FindCustommer(username)
	if err != nil {
		helper.Exception(ctx, false, "User not found", err)
		return
	}
	is_login := helper.CheckPasswordHash(password, customer.Password)
	if !is_login {
		helper.Exception(ctx, false, "Password or username is incorrect", err)
		return
	}

	randomToken := make([]byte, 32)
	_, errRand := rand.Read(randomToken)

	if errRand != nil {
		helper.Exception(ctx, false, "token error", err)
		return
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)

	token_request := customer_auth.CustomerTokenRequest{
		CustomerId: uint(customer.ID),
		AuthToken:  authToken,
	}

	data_token, err := c.service.GeneratedToken(token_request)
	if err != nil {
		helper.Exception(ctx, false, "token error", err)
		return
	}
	helper.Responses(ctx, false, "Success", data_token, 0, 0)
}

func (c *CustomerAuthController) Index(ctx *gin.Context) {
	tokens, err := c.service.FindAll()
	if err != nil {
		helper.Exception(ctx, false, "token error", err)
		return
	}
	helper.Responses(ctx, false, "Success", tokens, 0, 0)
}
