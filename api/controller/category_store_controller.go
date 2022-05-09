package controller

import (
	"api/helper"
	"api/middleware"
	category_store "api/modules/category/store"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryStoreController struct {
	service category_store.Service
}

func NewCategoryStoreController(service category_store.Service) *CategoryStoreController {
	return &CategoryStoreController{service}
}

func (c *CategoryStoreController) Index(ctx *gin.Context) {
	entities := ctx.Query("entities")
	categories, err := c.service.FindAll(entities, *middleware.StoreId)
	if err != nil {
		helper.Exception(ctx, false, "Category not found", err)
		return
	}

	var category_responses []category_store.Response

	for _, v := range categories {
		category_responses = append(category_responses, category_store.Response{
			Id:          uint(v.ID),
			Name:        v.Name,
			Slug:        v.Slug,
			Description: v.Description,
			StoreId:     v.StoreId,
			UserId:      v.UserId,
			User:        v.User,
			Store:       v.Store,
		})
	}

	helper.Responses(ctx, true, "Success", category_responses, 0, 0)
}

func (c *CategoryStoreController) Created(ctx *gin.Context) {
	categories, err := c.service.FindAll("", *middleware.StoreId)
	reqeust := category_store.Request{
		Name:        ctx.PostForm("name"),
		Slug:        helper.GeneratedCategorySlug(ctx.PostForm("name"), categories),
		Description: ctx.PostForm("description"),
		StoreId:     *middleware.StoreId,
		UserId:      middleware.UserId,
	}
	category, err := c.service.Created(reqeust)
	if err != nil {
		helper.Exception(ctx, false, "Category not created", err)
		return
	}
	helper.Responses(ctx, true, "Success", category, 0, 0)
}

func (c *CategoryStoreController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	entities := ctx.Query("entities")
	category, err := c.service.Show(uint(id), entities)
	if err != nil || category.ID == 0 {
		helper.Exception(ctx, false, "Category not found", err)
		return
	}
	category_responses := category_store.Response{
		Id:          uint(category.ID),
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
		StoreId:     category.StoreId,
		UserId:      category.UserId,
		User:        category.User,
		Store:       category.Store,
	}
	helper.Responses(ctx, true, "Success", category_responses, 0, 0)
}

func (c *CategoryStoreController) Updated(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	check, err := c.service.Show(uint(id), "")
	var category category_store.Request
	categories, err := c.service.FindAll("", *middleware.StoreId)
	if err != nil || check.ID == 0 {
		helper.Exception(ctx, false, "Category not found", err)
		return
	}
	category.Name = ctx.PostForm("name")
	category.Slug = helper.GeneratedCategorySlug(ctx.PostForm("name"), categories)
	category.Description = ctx.PostForm("description")

	category_upadate, err := c.service.Updated(uint(id), category)
	if err != nil {
		helper.Exception(ctx, false, "Category not updated", err)
		return
	}
	helper.Responses(ctx, true, "Success", category_upadate, 0, 0)
}

func (c *CategoryStoreController) Deleted(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	categories, err := c.service.Deleted(uint(id))
	if err != nil {
		helper.Exception(ctx, false, "Category not found", err)
		return
	}

	helper.Responses(ctx, true, "Success", categories, 0, 0)
}
