package customer

type CustomerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
	Contact  string `json:"contact"`
	StoreId  uint   `json:"store_id"`
}
