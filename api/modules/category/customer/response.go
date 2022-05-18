package category_customer

import "api/entity"

type Response struct {
	Id          uint             `json:"id"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	UserId      *uint            `json:"user_id"`
	StoreId     uint             `json:"store_id"`
	User        entity.User      `json:"user"`
	Store       entity.Store     `json:"store"`
	Product     []entity.Product `json:"product"`
}
