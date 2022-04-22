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
	FindAll(q string, page int, page_size int) ([]entity.TShirt, error)
	FindById(id int) (entity.TShirt, error)
	Created(reqeuest TshirtRequest) (entity.TShirt, error)
	Updated(request TshirtRequest, id int) (entity.TShirt, error)
	Deleted(id int) (entity.TShirt, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAll(q string, page int, page_size int) ([]entity.TShirt, error) {
	return service.repository.FindAll(q, page, page_size)
}

func (service *service) Created(request TshirtRequest) (entity.TShirt, error) {
	var tshirt_request entity.TShirt
	data, err := service.repository.FindAll("", 0, 0)
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

func (service *service) FindById(id int) (entity.TShirt, error) {
	return service.repository.FindById(id)
}

func (service *service) Updated(request TshirtRequest, id int) (entity.TShirt, error) {
	data, err := service.repository.FindAll("", 0, 0)

	dataIndex, _ := service.repository.FindById(id)

	dataIndex.Name = request.Name
	dataIndex.Slug = helper.GeneratedSlug(request.Name, data)
	dataIndex.Price = request.Price
	dataIndex.Stock = request.Stock
	dataIndex.Description = request.Description
	dataIndex.Image = request.Image
	dataIndex.UserId = uint(middleware.UserId)

	tshirt, err := service.repository.Updated(dataIndex)

	return tshirt, err
}

func (service *service) Deleted(id int) (entity.TShirt, error) {
	data, err := service.repository.FindById(id)
	if err != nil {
		return data, err
	}
	return service.repository.Delete(data)
}
