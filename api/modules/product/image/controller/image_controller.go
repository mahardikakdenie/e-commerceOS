package controller

import (
	"api/helper"
	"api/modules/product/image"
	rr "api/modules/product/image/RR"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	service image.Service
}

func NewController(service image.Service) *ImageController {
	return &ImageController{service}
}

func (c *ImageController) Index(ctx *gin.Context) {
	product_id, _ := strconv.Atoi(ctx.Param("product_id"))
	data, err := c.service.GetImageProduct(ctx, uint(product_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var imageResponses []rr.ProductImageResponse
	for _, v := range data {
		imageResponses = append(imageResponses, rr.ProductImageResponse{
			ID:        uint(v.ID),
			Image:     v.Image,
			ProductId: uint(v.ProductId),
			StoreId:   uint(v.StoreId),
			UserId:    uint(*v.UserId),
			Product:   v.Product,
			User:      v.User,
			Store:     v.Store,
		})
	}

	helper.Responses(ctx, true, "success", imageResponses, 0, 0)
}

func (c *ImageController) Store(ctx *gin.Context) {
	var request rr.ProductImageRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.service.InsertImage(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageResponses := rr.ProductImageResponse{
		ID:        uint(data.ID),
		Image:     data.Image,
		ProductId: uint(data.ProductId),
		StoreId:   uint(data.StoreId),
		UserId:    uint(*data.UserId),
		Product:   data.Product,
		User:      data.User,
		Store:     data.Store,
	}

	helper.Responses(ctx, true, "success", imageResponses, 0, 0)
}
