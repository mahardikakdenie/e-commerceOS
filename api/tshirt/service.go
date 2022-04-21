package tshirt

import (
	"api/entity"
	"api/helper"
	"api/middleware"
)

type service struct {
	repository Repository
}

type Service interface {
	FindAll() ([]entity.TShirt, error)
	Created(reqeuest TshirtRequest) (entity.TShirt, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAll() ([]entity.TShirt, error) {
	return service.repository.FindAll()
}

func (service *service) Created(request TshirtRequest) (entity.TShirt, error) {
	var tshirt_request entity.TShirt
	data, err := service.repository.FindAll()
	tshirt_request = entity.TShirt{
		Name:        request.Name,
		Slug:        helper.GeneratedSlug(request.Name, data),
		Price:       request.Price,
		Stock:       request.Stock,
		Description: request.Description,
		Image:       request.Image,
		UserId:      uint(middleware.UserId),
	}
	tshirt, err := service.repository.Created(tshirt_request)

	return tshirt, err
}
