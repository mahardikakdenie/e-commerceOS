package cart_customer

import "api/entity"

type CartRequest struct {
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
	StoreId    uint   `json:"store_id"`
	ProductId  uint   `json:"product_id"`
	CustomerId uint   `json:"customer_id"`
}

type CartResponses struct {
	Id         uint            `json:"id"`
	Quantity   int             `json:"quantity"`
	Status     string          `json:"status"`
	StoreId    uint            `json:"store_id"`
	ProductId  uint            `json:"product_id"`
	CustomerId uint            `json:"customer_id"`
	Store      entity.Store    `json:"store"`
	Customer   entity.Customer `json:"customer"`
	Product    entity.Product  `json:"product"`
}
