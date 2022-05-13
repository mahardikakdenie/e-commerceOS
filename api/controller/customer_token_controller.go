package controller

import (
	"api/helper"
	"api/middleware"
	"api/modules/customer_auth"
	"crypto/rand"
	"encoding/base64"
	"fmt"

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
	store_slug := ctx.PostForm("store_slug")
	if username == "" || password == "" || store_slug == "" {
		helper.Exception(ctx, false, "Record Tidak ditemukan", nil)
		return
	}

	customer, customer_err := c.service.FindCustommer(username)
	if customer_err != nil {
		helper.Exception(ctx, false, "customer not found in record", customer_err)
		return
	}
	store, store_err := c.service.FindStoreBySlug(store_slug)

	if store_err != nil {
		helper.Exception(ctx, false, "store not found", store_err)
		return
	}
	fmt.Println("customer: ", username)
	fmt.Println("store => ", store.ID != customer.StoreId)
	if store.ID != customer.StoreId || customer.StoreId == 0 {
		helper.Exception(ctx, false, "Account not found", store_err)
		return
	}
	is_login := helper.CheckPasswordHash(password, customer.Password)
	if !is_login {
		helper.Exception(ctx, false, "Password or username is incorrect", customer_err)
		return
	}

	// fmt.Println("Password => ", password)
	randomToken := make([]byte, 32)
	_, errRand := rand.Read(randomToken)

	if errRand != nil {
		helper.Exception(ctx, false, "token error", errRand)
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
	responses := customer_auth.CustomerTokenResponses{
		Token:      data_token.AuthToken,
		CustomerId: data_token.CustomerId,
		Customer:   data_token.Customer,
	}
	helper.Responses(ctx, true, "Success", responses, 0, 0)
}

func (c *CustomerAuthController) Index(ctx *gin.Context) {
	tokens, err := c.service.FindAll()
	if err != nil {
		helper.Exception(ctx, false, "token error", err)
		return
	}
	helper.Responses(ctx, false, "Success", tokens, 0, 0)
}

func (c *CustomerAuthController) Logout(ctx *gin.Context) {
	data, err := c.service.Logout(middleware.Token)
	if err != nil {
		helper.Exception(ctx, false, "token error", err)
		return
	}
	helper.Responses(ctx, true, "Success", data, 0, 0)
}
