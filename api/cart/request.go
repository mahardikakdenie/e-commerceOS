package cart

type CartRequest struct {
	UserID   uint   `json:"user_id"`
	TShirtID uint   `json:"tshirt_id"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}
