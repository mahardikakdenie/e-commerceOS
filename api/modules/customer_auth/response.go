package customer_auth

import "api/entity"

type CustomerTokenResponses struct {
	Token      string          `json:"token"`
	CustomerId uint            `json:"customer_id"`
	Customer   entity.Customer `json:"customer"`
}
