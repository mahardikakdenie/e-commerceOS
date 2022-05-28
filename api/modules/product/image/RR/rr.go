package rr

import "api/entity"

type ProductImageRequest struct {
	ProductID uint   `json:"product_id"`
	Image     string `json:"image"`
	StoreId   uint   `json:"store_id"`
	UserId    uint   `json:"user_id"`
}

type ProductImageResponse struct {
	ID        uint           `json:"id"`
	ProductId uint           `json:"product_id"`
	Image     string         `json:"url"`
	StoreId   uint           `json:"store_id"`
	Product   entity.Product `json:"product"`
	Store     entity.Store   `json:"store"`
	UserId    uint           `json:"user_id"`
	User      entity.User    `json:"user"`
}
