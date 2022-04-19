package auth

import (
	"api/entity"
	"api/user"
)

type Service interface {
	Login(Request AuthRequest) (entity.Auth, error)
	FindByEmail(email string) (entity.User, error)
	Register(request user.UserRequest) (entity.User, error)
	FindByToken(token string) (entity.Auth, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Login(Request AuthRequest) (entity.Auth, error) {
	var auth entity.Auth

	auth = entity.Auth{
		UserId:    Request.UserId,
		AuthToken: Request.AuthToken,
	}

	return s.repository.Login(auth)
}

func (s *service) FindByEmail(email string) (entity.User, error) {
	return s.repository.FindByEmail(email)
}

func (s *service) Register(request user.UserRequest) (entity.User, error) {
	var user entity.User
	user = entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	return s.repository.Register(user)
}

func (s *service) FindByToken(token string) (entity.Auth, error) {
	return s.repository.FindByToken(token)
}
