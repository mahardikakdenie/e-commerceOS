package controller

import (
	"api/entity"
	"api/helper"
	"api/middleware"
	cart_customer "api/modules/cart/customer"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartsController struct {
	service cart_customer.CartCussServiceInterface
}

func NewCartsController(service cart_customer.CartCussServiceInterface) *CartsController {
	return &CartsController{
		service: service,
	}
}

func (c *CartsController) Index(ctx *gin.Context) {
	entities := ctx.Query("entities")
	customer_id, _ := strconv.Atoi(ctx.Param("customer_id"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	per_page, _ := strconv.Atoi(ctx.Query("per_page"))

	carts, err, customer_err := c.service.IndexCartCustomer(entities, uint(customer_id), page, per_page)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}
	if customer_err != nil {
		helper.Exception(ctx, false, customer_err.Error(), nil)
		return
	}

	helper.Responses(ctx, true, "Success", carts, page, per_page)
}

func (c *CartsController) Store(ctx *gin.Context) {
	carts := cart_customer.CartRequest{}
	if err := ctx.ShouldBind(&carts); err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}

	store_id, _ := strconv.Atoi(ctx.Param("store_id"))
	product_id, _ := strconv.Atoi(ctx.Param("product_id"))

	data := entity.Cart{
		Quantity:   carts.Quantity,
		CustomerId: middleware.CustomerId,
		StoreID:    uint(store_id),
		ProductID:  uint(product_id),
		Status:     carts.Status,
	}

	if cart, err := c.service.InsertCart(data, uint(store_id), uint(product_id)); err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	} else {
		helper.Responses(ctx, true, "Success", cart, 0, 0)
	}
}
