package controller

import (
	"api/helper"
	"api/middleware"
	product_store "api/modules/product/store"
	"fmt"
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
	store_slug := ctx.Query("store_slug")
	fmt.Println("Store_slug := ", store_slug)
	products, err := c.service.FindAll(entities, store_slug)
	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}

	var response []product_store.ProductStoreResponse

	for _, v := range products {
		response = append(response, product_store.ProductStoreResponse{
			Name:       v.Name,
			Price:      int(v.Price),
			Stock:      int(v.Stock),
			Image:      v.Image,
			CategoryId: uint(v.CategoryId),
			StoreId:    uint(v.StoreId),
			UserId:     uint(*v.UserId),
			User:       v.User,
			Category:   v.Category,
			Store:      v.Store,
		})
	}

	helper.Responses(ctx, true, "Success", response, 0, 0)
}

func (c *ProductStoreController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	store_id := ctx.Query("store_id")
	compare_id, _ := strconv.Atoi(store_id)
	entities := ctx.Query("entities")
	product, err := c.service.FindById(uint(id), store_id, entities)

	if err != nil {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	if uint(compare_id) != product.StoreId {
		helper.Exception(ctx, false, "Product not found", err)
		return
	}
	response := product_store.ProductStoreResponse{
		Name:       product.Name,
		Price:      int(product.Price),
		Stock:      int(product.Stock),
		Image:      product.Image,
		CategoryId: uint(product.CategoryId),
		StoreId:    uint(product.StoreId),
		UserId:     uint(*product.UserId),
		User:       product.User,
		Category:   product.Category,
		Store:      product.Store,
	}
	helper.Responses(ctx, true, "Success", response, 0, 0)
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
