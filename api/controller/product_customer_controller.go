package controller

import (
	"api/helper"
	product_customer "api/modules/product/customer"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCustomerController struct {
	service product_customer.Service
}

func NewProductCustomerController(service product_customer.Service) *ProductCustomerController {
	return &ProductCustomerController{service}
}

func (c *ProductCustomerController) Index(ctx *gin.Context) {
	q := ctx.Query("q")
	entities := ctx.Query("entities")
	page, _ := strconv.Atoi(ctx.Query("page"))
	per_page, _ := strconv.Atoi(ctx.Query("per_page"))
	category_id := ctx.Query("category_id")
	is_seller, _ := strconv.Atoi(ctx.Query("is_seller"))
	store_id := ctx.Query("store_id")
	products, err := c.service.FindAll(q, entities, page, per_page, category_id, is_seller, store_id)
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	var data []product_customer.ProductResponse
	for _, v := range products {
		data = append(data, product_customer.ProductResponse{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			View:        v.View,
			CategoryId:  v.CategoryId,
			UserId:      v.UserId,
			StoreId:     v.StoreId,
			Entity: product_customer.Entity{
				Category: v.Category,
				User:     v.User,
				Store:    v.Store,
			},
			Link: product_customer.Link{
				Image: v.Image,
			},
		})
	}

	helper.Responses(ctx, true, "Success", data, page, per_page)
}

func (c *ProductCustomerController) ProductByCategory(ctx *gin.Context) {
	slug := ctx.Param("slug")
	page, _ := strconv.Atoi(ctx.Query("page"))
	per_page, _ := strconv.Atoi(ctx.Query("per_page"))
	entities := ctx.Query("entities")
	is_seller, _ := strconv.Atoi(ctx.Query("is_seller"))
	products, err, category_err := c.service.FindByCategorySlug(slug, entities, page, per_page, is_seller, ctx.Query("store_id"))
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	if category_err != nil {
		helper.Exception(ctx, false, "Category not found", category_err)
		return
	}
	var data []product_customer.ProductResponse
	for _, v := range products {
		data = append(data, product_customer.ProductResponse{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			View:        v.View,
			CategoryId:  v.CategoryId,
			UserId:      v.UserId,
			StoreId:     v.StoreId,
			Entity: product_customer.Entity{
				Category: v.Category,
				User:     v.User,
				Store:    v.Store,
			},
			Link: product_customer.Link{
				Image: v.Image,
			},
		})
	}
	helper.Responses(ctx, true, "Success", data, page, per_page)
}

func (c *ProductCustomerController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	entities := ctx.Query("entities")
	category_id := ctx.Query("category_id")
	store_id := ctx.Query("store_id")
	product, err := c.service.FindById(uint(id), entities, category_id, store_id)
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	helper.Responses(ctx, true, "Success", product_customer.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		View:        product.View,
		CategoryId:  product.CategoryId,
		UserId:      product.UserId,
		StoreId:     product.StoreId,
		Entity: product_customer.Entity{
			Category: product.Category,
			User:     product.User,
			Store:    product.Store,
		},
		Link: product_customer.Link{
			Image: product.Image,
		},
	}, 1, 1)
}
