package cart

type CartResponses struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	TShirtId int    `json:"tshirt_id"`
	Qty      int    `json:"quantity"`
	Status   string `json:"status"`
}
