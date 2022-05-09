package helper

import (
	"api/entity"
	"api/modules/user"
)

func UserResponses(v entity.User) user.UserResponses {
	var user_responses user.UserResponses

	user_responses = user.UserResponses{
		ID:    uint(v.ID),
		Name:  v.Name,
		Email: v.Email,
		Auth:  v.Auth,
		Role:  v.Role,
	}

	return user_responses
}
