package cart_customer

import "api/entity"

type CartRequest struct {
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
	StoreId    uint   `json:"store_id"`
	ProductId  uint   `json:"product_id"`
	CustomerId uint   `json:"customer_id"`
	Variant    string `json:"variant"`
}

type CartResponses struct {
	Id         uint            `json:"id"`
	Quantity   int             `json:"quantity"`
	Status     string          `json:"status"`
	StoreId    uint            `json:"store_id"`
	ProductId  uint            `json:"product_id"`
	CustomerId uint            `json:"customer_id"`
	Variant    string          `json:"variant"`
	Store      entity.Store    `json:"store"`
	Customer   entity.Customer `json:"customer"`
	Product    Product         `json:"product"`
}

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Image       string `json:"image"`
}
