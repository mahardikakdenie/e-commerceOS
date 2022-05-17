package product_customer

import "api/entity"

type ProductResponse struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	View        int     `json:"view"`
	CategoryId  uint    `json:"category_id"`
	UserId      *uint   `json:"user_id"`
	StoreId     uint    `json:"store_id"`
	Entity      Entity  `json:"entities"`
	Link        Link    `json:"link"`
}

type Entity struct {
	Category entity.Category `json:"category"`
	User     entity.User     `json:"user"`
	Store    entity.Store    `json:"store"`
}

type Link struct {
	Image string `json:"image"`
}
