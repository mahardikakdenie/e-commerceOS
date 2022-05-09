package controller

import (
	"api/helper"
	"api/middleware"
	"api/modules/cart"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	service cart.Service
}

func NewCartController(service cart.Service) *CartController {
	return &CartController{service}
}

func (c *CartController) Index(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	per_page, _ := strconv.Atoi(ctx.Query("per_page"))
	entities := ctx.Query("entities")
	data, err := c.service.FindAll(entities)
	if err != nil {
		helper.Exception(ctx, false, "Cart not found", nil)
		return
	}

	var cartResponses []cart.CartResponses
	for _, v := range data {
		cartResponses = append(cartResponses, cart.CartResponses{
			ID:     int(v.ID),
			UserId: uint(v.UserID),
			// TShirtId: int(v.TShirtID),
			Qty:    int(v.Quantity),
			Status: v.Status,
			User:   v.User,
			// TShirt:   v.TShirt,
		})
	}

	helper.Responses(ctx, true, "Success", cartResponses, page, per_page)
}

func (c *CartController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, "Cart not found", nil)
		return
	}

	var cartResponses cart.CartResponses
	cartResponses = cart.CartResponses{
		ID:     int(data.ID),
		UserId: uint(data.UserID),
		// TShirtId: int(data.TShirtID),
		Qty:    int(data.Quantity),
		Status: data.Status,
	}

	helper.Responses(ctx, true, "Success", cartResponses, 0, 0)
}

func (c *CartController) Store(ctx *gin.Context) {
	var cart cart.CartRequest
	err := ctx.BindJSON(&cart)
	if err != nil {
		helper.Exception(ctx, false, "Invalid data", nil)
		return
	}

	data, err := c.service.Created(cart)
	if err != nil {
		helper.Exception(ctx, false, "Failed to store cart", nil)
		return
	}

	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *CartController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	cartId, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, "Cart not found", nil)
		return
	}
	// tshirt_id, _ := strconv.Atoi(ctx.PostForm("tshirt_id"))
	quantity, _ := strconv.Atoi(ctx.PostForm("quantity"))
	var request cart.CartRequest
	request = cart.CartRequest{
		// TShirtID: helper.CheckerCartTshirtId(uint(tshirt_id), cartId),
		UserID:   uint(middleware.UserId),
		Quantity: helper.CheckCartQuantity(quantity, cartId),
		Status:   helper.CheckCartStatus(ctx.PostForm("status"), cartId),
	}

	data, err := c.service.Updated(request, id)
	if err != nil {
		helper.Exception(ctx, false, "Failed to update cart", nil)
		return
	}

	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *CartController) Destroy(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := c.service.Deleted(id)
	if err != nil {
		helper.Exception(ctx, false, "Failed to delete cart", nil)
		return
	}

	helper.Responses(ctx, true, "Success", data, 0, 0)
}
