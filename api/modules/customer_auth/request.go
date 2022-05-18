package customer_auth

type CustomerTokenRequest struct {
	CustomerId uint   `json:"customer_id"`
	AuthToken  string `json:"auth_token"`
}

type CustomerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
	StoreId  uint   `json:"store_id"`
}
