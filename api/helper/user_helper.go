package helper

import (
	"api/entity"
	"api/user"
)

func UserResponses(v entity.User) user.UserReponse {
	var user_responses user.UserReponse

	user_responses = user.UserReponse{
		ID:    int(v.ID),
		Name:  v.Name,
		Email: v.Email,
		Auth:  v.Auth,
		Role:  v.Role,
	}

	return user_responses
}
