package orders

import "api/entity"

type OrderResponse struct {
	ID       uint          `json:"id"`
	UserID   uint          `json:"user_id"`
	TShirtID uint          `json:"tshirt_id"`
	Quantity int           `json:"quantity"`
	Status   string        `json:"status"`
	Address  string        `json:"address"`
	User     entity.User   `json:"user"`
	TShirt   entity.TShirt `json:"tshirt"`
}
