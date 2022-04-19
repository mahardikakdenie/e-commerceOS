package user

import "api/entity"

type Service interface {
	FindAll() (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Me(user_id int) (entity.User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAll() (entity.User, error) {
	return service.repository.FindAll()
}

func (service *service) FindByEmail(email string) (entity.User, error) {
	return service.repository.FindByEmail(email)
}

func (service *service) Me(user_id int) (entity.User, error) {
	return service.repository.Me(user_id)
}
