package order

type OrderRequest struct {
	UserID     uint   `json:"user_id"`
	TShirtID   uint   `json:"tshirt_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
	Address    string `json:"address"`
	CustomerID uint   `json:"customer_id"`
}
