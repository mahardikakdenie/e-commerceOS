package user

import (
	"api/entity"
)

type UserReponse struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Role  entity.Role   `json:"role"`
	Auth  []entity.Auth `json:"auth"`
}
