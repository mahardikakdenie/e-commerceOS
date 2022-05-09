package role

import "api/entity"

type Service interface {
	FindAll() ([]entity.Role, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAll() ([]entity.Role, error) {
	return service.repository.FindAll()
}
