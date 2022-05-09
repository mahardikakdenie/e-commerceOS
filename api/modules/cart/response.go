package cart

import "api/entity"

type CartResponses struct {
	ID       int           `json:"id"`
	UserId   uint          `json:"user_id"`
	TShirtId int           `json:"tshirt_id"`
	Qty      int           `json:"quantity"`
	Status   string        `json:"status"`
	User     entity.User   `json:"user"`
	TShirt   entity.TShirt `json:"tshirt"`
}
