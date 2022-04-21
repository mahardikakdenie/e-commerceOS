package controller

import (
	"api/tshirt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TShirtController struct {
	service tshirt.Service
}

func NewTShirtController(service tshirt.Service) *TShirtController {
	return &TShirtController{service}
}

func (controller *TShirtController) FindAll(ctx *gin.Context) {
	tshirts, err := controller.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": err.Error(),
			},
			"data": gin.H{},
		})
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
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
		"data": resposes,
	})
}

func (controler *TShirtController) Store(ctx *gin.Context) {
	var tshirt tshirt.TshirtRequest
	ctx.ShouldBindJSON(&tshirt)
	data, err := controler.service.Created(tshirt)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": err.Error(),
			},
			"data": gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{},
		"data": data,
	})
}
