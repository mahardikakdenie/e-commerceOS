package product_store

import "api/entity"

type ProductStoreResponse struct {
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Stock       int             `json:"stock"`
	Image       string          `json:"image"`
	CategoryId  uint            `json:"category_id"`
	StoreId     uint            `json:"store_id"`
	UserId      uint            `json:"user_id"`
	Store       entity.Store    `json:"store"`
	Category    entity.Category `json:"category"`
	User        entity.User     `json:"user"`
}
