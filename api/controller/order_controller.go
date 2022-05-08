package controller

import (
	"api/helper"
	"api/middleware"
	"api/order"
	orders "api/order"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service orders.Service
}

func NewOrderController(service orders.Service) *OrderController {
	return &OrderController{service}
}

func (c *OrderController) Index(ctx *gin.Context) {
	q := ctx.Query("q")
	page_string := ctx.Query("page")
	page, _ := strconv.Atoi(page_string)
	page_size_string := ctx.Query("per_page")
	per_page, _ := strconv.Atoi(page_size_string)
	entities := ctx.Query("entities")
	tshirts, err := c.service.FindAll(q, page, per_page, entities)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}

	var orderResponses []orders.OrderResponse
	for _, v := range tshirts {
		orderResponses = append(orderResponses, orders.OrderResponse{
			ID:       uint(v.ID),
			UserID:   &v.User.ID,
			TShirtID: uint(v.TShirtID),
			Quantity: int(v.Quantity),
			Status:   v.Status,
			Address:  v.Address,
			User:     v.User,
			TShirt:   v.TShirt,
		})
	}

	helper.Responses(ctx, true, "Success", orderResponses, page, per_page)
}

func (c *OrderController) Show(ctx *gin.Context) {
	id_string := ctx.Param("id")
	id, _ := strconv.Atoi(id_string)
	data, err := c.service.FindById(int(id))
	var orderResponse orders.OrderResponse
	orderResponse = orders.OrderResponse{
		ID:       uint(data.ID),
		UserID:   data.UserID,
		TShirtID: uint(data.TShirtID),
		Quantity: int(data.Quantity),
		Status:   data.Status,
		Address:  data.Address,
		User:     data.User,
		TShirt:   data.TShirt,
	}
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}
	helper.Responses(ctx, true, "Success", orderResponse, 0, 0)
}

func (c *OrderController) Store(ctx *gin.Context) {
	var request orders.OrderRequest
	tshirt_id_string := ctx.PostForm("tshirt_id")
	tshirt_id, _ := strconv.Atoi(tshirt_id_string)
	quantity, _ := strconv.Atoi(ctx.PostForm("quantity"))
	customer_id, _ := strconv.Atoi(ctx.PostForm("customer_id"))

	request = orders.OrderRequest{
		UserID:     uint(middleware.UserId),
		Status:     "pending",
		TShirtID:   uint(tshirt_id),
		Address:    ctx.PostForm("address"),
		Quantity:   quantity,
		CustomerID: uint(customer_id),
	}

	data, err := c.service.Created(request)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}
	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *OrderController) Update(ctx *gin.Context) {
	id_string := ctx.Param("id")
	id, _ := strconv.Atoi(id_string)
	dataOrder, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}
	tshirt_id, _ := strconv.Atoi(ctx.PostForm("tshirt_id"))
	quantity, _ := strconv.Atoi(ctx.PostForm("quantity"))
	var request order.OrderRequest
	request = order.OrderRequest{
		Status:   helper.CheckStatus(ctx.PostForm("status"), dataOrder),
		Address:  helper.CheckAddress(ctx.PostForm("address"), dataOrder),
		TShirtID: helper.CheckerTshirtId(uint(tshirt_id), dataOrder),
		Quantity: helper.CheckQuantity(quantity, dataOrder),
	}
	data, err := c.service.Updated(id, request)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}
	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *OrderController) Destroy(ctx *gin.Context) {
	id_string := ctx.Param("id")
	id, _ := strconv.Atoi(id_string)
	order, err := c.service.Deleted(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), nil)
		return
	}
	helper.Responses(ctx, true, "Success", order, 0, 0)
}
