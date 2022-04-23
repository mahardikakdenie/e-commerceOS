package cart

import (
	"api/entity"
	"api/middleware"
)

type Service interface {
	FindAll() ([]entity.Cart, error)
	Created(request CartRequest) (entity.Cart, error)
	FindById(id int) (entity.Cart, error)
	Deleted(id int) (entity.Cart, error)
	Updated(request CartRequest, id int) (entity.Cart, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAll() ([]entity.Cart, error) {
	return service.repository.FindAll()
}

func (service *service) Created(request CartRequest) (entity.Cart, error) {
	var cart entity.Cart
	cart = entity.Cart{
		UserID:   uint(middleware.UserId),
		TShirtID: request.TShirtID,
		Quantity: request.Quantity,
		Status:   "pending",
	}
	carts, err := service.repository.Created(cart)
	return carts, err
}

func (s *service) FindById(id int) (entity.Cart, error) {
	cart, err := s.repository.FindById(id)
	return cart, err
}

func (s *service) Deleted(id int) (entity.Cart, error) {
	cart, err := s.repository.FindById(id)
	delete_cart, err := s.repository.Deleted(cart)
	return delete_cart, err
}

func (s *service) Updated(request CartRequest, id int) (entity.Cart, error) {
	cart, err := s.repository.FindById(id)
	cart = entity.Cart{
		UserID:   uint(middleware.UserId),
		TShirtID: request.TShirtID,
		Quantity: request.Quantity,
		Status:   request.Status,
	}
	carts, err := s.repository.Updated(cart)
	return carts, err
}
