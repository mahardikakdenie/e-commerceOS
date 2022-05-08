package customer_auth

type CustomerTokenRequest struct {
	CustomerId uint   `json:"customer_id"`
	AuthToken  string `json:"auth_token"`
}
