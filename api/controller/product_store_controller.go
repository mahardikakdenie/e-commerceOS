package controller

import (
	"api/helper"
	"api/middleware"
	product_store "api/modules/product/store"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductStoreController struct {
	service product_store.Service
}

func NewProductStoreController(service product_store.Service) *ProductStoreController {
	return &ProductStoreController{service}
}

func (c *ProductStoreController) Index(ctx *gin.Context) {
	entities := ctx.Query("entities")
	products, err := c.service.FindAll(entities)
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}

	helper.Responses(ctx, true, "Success", products, 0, 0)
}

func (c *ProductStoreController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.service.FindById(uint(id))
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	helper.Responses(ctx, true, "Success", product, 0, 0)
}

func (c *ProductStoreController) Created(ctx *gin.Context) {
	price, _ := strconv.Atoi(ctx.PostForm("price"))
	stock, _ := strconv.Atoi(ctx.PostForm("stock"))
	category_id, _ := strconv.Atoi(ctx.PostForm("category_id"))
	var request product_store.ProductStoreRequest
	request = product_store.ProductStoreRequest{
		Name:        ctx.PostForm("name"),
		Description: ctx.PostForm("description"),
		Price:       int(price),
		Stock:       int(stock),
		Image:       ctx.PostForm("image"),
		CategoryId:  uint(category_id),
		StoreId:     *middleware.StoreId,
		UserId:      middleware.UserId,
	}

	data, err := c.service.Created(request)
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *ProductStoreController) Updated(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	price, _ := strconv.Atoi(ctx.PostForm("price"))
	stock, _ := strconv.Atoi(ctx.PostForm("stock"))
	category_id, _ := strconv.Atoi(ctx.PostForm("category_id"))
	var request product_store.ProductStoreRequest
	request = product_store.ProductStoreRequest{
		Name:        ctx.PostForm("name"),
		Description: ctx.PostForm("description"),
		Price:       int(price),
		Stock:       int(stock),
		Image:       ctx.PostForm("image"),
		CategoryId:  uint(category_id),
		StoreId:     *middleware.StoreId,
		UserId:      middleware.UserId,
	}

	product, err := c.service.Updated(request, uint(id))
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	helper.Responses(ctx, true, "Success", product, 0, 0)
}

func (c *ProductStoreController) Deleted(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.service.Deleted(uint(id))
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	helper.Responses(ctx, true, "Success", product, 0, 0)
}
