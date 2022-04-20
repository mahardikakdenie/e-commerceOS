package controller

import (
	"api/entity"
	"api/user"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func userResponses(v entity.User) user.UserReponse {
	var user_responses user.UserReponse

	user_responses = user.UserReponse{
		ID:    int(v.ID),
		Name:  v.Name,
		Email: v.Email,
		Auth:  v.Auth,
	}

	return user_responses
}
