package cart_customer

import "api/entity"

type Service interface {
	FindAll(q string, entities string, page int, per_page int) (carts []entity.Cart, err error, customer_err error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(q string, entities string, page int, per_page int) (carts []entity.Cart, err error, customer_err error) {
	return s.repository.
		FindAll(q, entities, page, per_page)
}
