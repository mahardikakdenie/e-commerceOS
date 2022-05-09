package controller

import (
	"api/entity"
	"api/helper"
	"api/modules/store"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	service store.Service
}

func NewStoreController(service store.Service) *StoreController {
	return &StoreController{service}
}

func (c *StoreController) FindAll(ctx *gin.Context) {
	q := ctx.Query("q")
	entities := ctx.Query("entities")
	stores, err := c.service.FindAll(q, entities)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	data_responses := DataStoreResponses(stores)
	helper.Responses(ctx, true, "Success", data_responses, 0, 0)
}

func (c *StoreController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	store, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}

	data := DataStoreResponse(store)

	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *StoreController) Created(ctx *gin.Context) {
	var request store.StoreRequest
	stores, err := c.service.FindAll("", "")
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	request = store.StoreRequest{
		Name:        ctx.PostForm("name"),
		Slug:        helper.GeneratedStoreSlug(ctx.PostForm("name"), stores),
		Description: ctx.PostForm("description"),
	}

	data, err := c.service.Created(request)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}

	helper.Responses(ctx, true, "Success", data, 0, 0)

}

func (c *StoreController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var dataStore store.StoreRequest
	storeUpdate, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
	}

	name := ctx.PostForm("name")
	desc := ctx.PostForm("description")

	dataStore = store.StoreRequest{
		Name:        helper.CheckStoreName(name, storeUpdate),
		Slug:        helper.CheckerStroreSlug(name, storeUpdate),
		Description: helper.CheckStoreDescription(desc, storeUpdate),
	}

	data, err := c.service.Update(dataStore, id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}

	helper.Responses(ctx, true, "Success", data, 0, 0)
}

func (c *StoreController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	store, err := c.service.FindById(id)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}
	store_delete, err := c.service.Deleted(store)
	if err != nil {
		helper.Exception(ctx, false, err.Error(), err)
		return
	}

	helper.Responses(ctx, true, "Success", store_delete, 0, 0)
}

func DataStoreResponse(data entity.Store) store.StoreResponses {
	return store.StoreResponses{
		Name:      data.Name,
		Slug:      data.Slug,
		Desc:      data.Description,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func DataStoreResponses(data []entity.Store) []store.StoreResponses {
	var responses []store.StoreResponses
	for _, v := range data {
		responses = append(responses, store.StoreResponses{
			Name:      v.Name,
			Slug:      v.Slug,
			Desc:      v.Description,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		})
	}
	return responses
}
