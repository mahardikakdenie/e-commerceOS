package controller

import (
	"api/customer"
	"api/entity"
	"api/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service customer.Service
}

func NewCustomerController(service customer.Service) *CustomerController {
	return &CustomerController{service}
}

func (c *CustomerController) Index(ctx *gin.Context) {
	q := ctx.Query("q")
	page, _ := strconv.Atoi(ctx.Query("page"))
	per_page, _ := strconv.Atoi(ctx.Query("per_page"))
	entities := ctx.Query("entities")
	customers, err := c.service.FindAll(q, page, per_page, entities)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	helper.Responses(ctx, true, "Success", customers, page, per_page)
}

func (c CustomerController) Store(ctx *gin.Context) {
	var request customer.CustomerRequest
	hasPassword, _ := helper.HashPassword(ctx.PostForm("password"))
	StoreID, _ := strconv.Atoi(ctx.PostForm("store_id"))
	request = customer.CustomerRequest{
		Username: ctx.PostForm("username"),
		Password: hasPassword,
		Email:    ctx.PostForm("email"),
		Contact:  ctx.PostForm("contact"),
		Image:    ctx.PostForm("image_url"),
		StoreId:  uint(StoreID),
	}
	data, err := c.service.Created(request)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	response := Response(data)
	helper.Responses(ctx, true, "Success", response, 0, 0)
}

func Response(v entity.Customer) customer.CustomerResponse {
	db_data := customer.CustomerResponse{
		Id:        v.ID,
		Username:  v.Username,
		Email:     v.Email,
		Contact:   v.Contact,
		Image:     v.Image,
		UpdateAt:  v.UpdatedAt,
		CreateAt:  v.CreatedAt,
		DeletedAt: v.DeletedAt,
	}

	return db_data
}
