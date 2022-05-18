package controller

import (
	"api/helper"
	category_customer "api/modules/category/customer"

	"github.com/gin-gonic/gin"
)

type CategoryCustomerController struct {
	service category_customer.Service
}

func NewCategoryCustomerController(service category_customer.Service) *CategoryCustomerController {
	return &CategoryCustomerController{service}
}

func (c *CategoryCustomerController) Index(ctx *gin.Context) {
	entities := ctx.Query("entities")
	store_id := ctx.Query("store_id")
	categories, err := c.service.FindCategoryByStore(store_id, entities)
	if err != nil {
		helper.Exception(ctx, false, "Category not found", err)
		return
	}

	var category_reponses []category_customer.Response

	for _, v := range categories {
		category_reponses = append(category_reponses, category_customer.Response{
			Id:          uint(v.ID),
			Name:        v.Name,
			Slug:        v.Slug,
			Description: v.Description,
			StoreId:     v.StoreId,
			UserId:      v.UserId,
			User:        v.User,
			Store:       v.Store,
			Product:     v.Product,
		})
	}

	helper.Responses(ctx, true, "Success", category_reponses, 0, 0)
}
