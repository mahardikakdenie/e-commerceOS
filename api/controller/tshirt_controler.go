package controller

import (
	"api/helper"
	"api/middleware"
	"api/tshirt"

	"strconv"

	"github.com/gin-gonic/gin"
)

type TShirtController struct {
	service tshirt.Service
}

func NewTShirtController(service tshirt.Service) *TShirtController {
	return &TShirtController{service}
}

func (controller *TShirtController) FindAll(ctx *gin.Context) {
	q := ctx.Request.FormValue("q")
	page_string := ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(page_string)
	page_size_string := ctx.Request.FormValue("page_size")
	page_size, _ := strconv.Atoi(page_size_string)
	tshirts, err := controller.service.FindAll(q, page, page_size)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), tshirts)
		return
	}

	var resposes []tshirt.TshirtResponse
	var link tshirt.Link
	for _, v := range tshirts {
		link = tshirt.Link{
			Image: v.Image,
		}
		resposes = append(resposes, tshirt.TshirtResponse{
			ID:          uint(v.ID),
			Name:        v.Name,
			Price:       v.Price,
			Slug:        v.Slug,
			Description: v.Description,
			Stock:       v.Stock,
			UserId:      v.UserId,
			User:        v.User,
			Link:        link,
		})
	}

	helper.Responses(ctx, true, "Success", resposes, page, page_size)
}

func (controler *TShirtController) Store(ctx *gin.Context) {
	var tshirt tshirt.TshirtRequest
	ctx.ShouldBindJSON(&tshirt)
	data, err := controler.service.Created(tshirt)

	if err != nil {
		helper.Exception(ctx, false, err.Error(), data)
		return
	}
	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *TShirtController) Updated(ctx *gin.Context) {
	id_string := ctx.Param("id")
	id, _ := strconv.Atoi(id_string)
	data, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), data)
		return
	}
	if data.ID == 0 {
		helper.Exception(ctx, false, "Data not found", data)
		return
	}

	name := ctx.PostForm("name")
	price_string := ctx.PostForm("price")
	price, _ := strconv.Atoi(price_string)
	slug := ctx.PostForm("slug")
	description := ctx.PostForm("description")
	stock_string := ctx.PostForm("stock")
	stock, _ := strconv.Atoi(stock_string)
	image := ctx.PostForm("image")

	var request tshirt.TshirtRequest

	request = tshirt.TshirtRequest{
		Name:        helper.CheckerName(name, data),
		Price:       helper.CheckerPrice(price, data),
		Slug:        helper.CheckerSlug(slug, data),
		Description: helper.CheckDescription(description, data),
		Stock:       helper.CheckStock(stock, data),
		Image:       helper.CheckImage(image, data),
		UserId:      middleware.UserId,
	}

	dataUpdate, err := c.service.Updated(request, id)

	if err != nil {
		helper.Exception(ctx, false, err.Error(), dataUpdate)
		return
	}

	helper.Responses(ctx, true, "Success", dataUpdate, 0, 0)

}

func (c *TShirtController) Show(ctx *gin.Context) {
	id_string := ctx.Param("id")
	id, _ := strconv.Atoi(id_string)
	data, err := c.service.FindById(id)
	if data.ID == 0 {
		helper.Exception(ctx, false, "Data not found", data)
		return
	}
	if err != nil {
		helper.Exception(ctx, false, err.Error(), data)
		return
	}
	var resposes tshirt.TshirtResponse
	resposes = tshirt.TshirtResponse{
		ID:          uint(data.ID),
		Name:        data.Name,
		Price:       data.Price,
		Slug:        data.Slug,
		Description: data.Description,
		Stock:       data.Stock,
		UserId:      data.UserId,
		User:        data.User,
		Link: tshirt.Link{
			Image: data.Image,
		},
	}
	helper.Responses(ctx, true, "Success", resposes, 0, 0)
}

func (c *TShirtController) Destroy(ctx *gin.Context) {
	id_string := ctx.Param("id")
	id, _ := strconv.Atoi(id_string)
	data, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), data)
		return
	}
	if data.ID == 0 {
		helper.Exception(ctx, false, "Data not found", data)
		return
	}
	dataDelete, err := c.service.Deleted(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), dataDelete)
		return
	}
	helper.Responses(ctx, true, "Success", dataDelete, 0, 0)
}
