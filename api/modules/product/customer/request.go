package product_customer

import "api/entity"

type ProductRequest struct {
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	Price       float64         `json:"price"`
	Image       string          `json:"image"`
	Description string          `json:"description"`
	Stock       int             `json:"stock"`
	View        int             `json:"view"`
	CategoryId  uint            `json:"category_id"`
	UserId      *uint           `json:"user_id"`
	Category    entity.Category `json:"category"`
	User        entity.User     `json:"user"`
	StoreId     uint            `json:"store_id"`
	Store       entity.Store    `json:"store"`
}
