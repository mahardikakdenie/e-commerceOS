package user

import (
	"api/entity"
)

type UserResponses struct {
	ID    uint          `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Role  entity.Role   `json:"role"`
	Auth  []entity.Auth `json:"auth"`
}
