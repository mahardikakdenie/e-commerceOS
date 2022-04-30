package order

import (
	"api/entity"
	// "api/middleware"
)

type Service interface {
	FindAll(q string, page int, par_page int, entities string) ([]entity.Order, error)
	FindById(id int) (entity.Order, error)
	Created(order OrderRequest) (entity.Order, error)
	Updated(id int, request OrderRequest) (entity.Order, error)
	Deleted(int) (entity.Order, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAll(q string, page int, par_page int, entities string) ([]entity.Order, error) {
	return service.repository.FindAll(q, page, par_page, entities)
}

func (service *service) FindById(id int) (entity.Order, error) {
	return service.repository.FindById(id)
}

func (s *service) Created(request OrderRequest) (entity.Order, error) {
	var order entity.Order
	order = entity.Order{
		// UserID:   uint(middleware.UserId),
		TShirtID: request.TShirtID,
		Quantity: request.Quantity,
		Status:   request.Status,
		Address:  request.Address,
	}
	orders, err := s.repository.Created(order)
	return orders, err
}

func (s *service) Updated(id int, request OrderRequest) (entity.Order, error) {
	order, err := s.repository.FindById(id)
	if err != nil {
	}
	order = entity.Order{
		// UserID:   uint(middleware.UserId),
		TShirtID: request.TShirtID,
		Quantity: request.Quantity,
		Status:   request.Status,
		Address:  request.Address,
	}

	orders, err := s.repository.Updated(order)
	return orders, err
}

func (s *service) Deleted(id int) (entity.Order, error) {
	order, err := s.repository.FindById(id)
	if err != nil {
	}
	orders, err := s.repository.Deleted(order)
	return orders, err
}
