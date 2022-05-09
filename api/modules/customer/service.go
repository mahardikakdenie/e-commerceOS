package customer

import "api/entity"

type Service interface {
	FindAll(q string, page int, per_page int, entities string) ([]entity.Customer, error)
	Created(request CustomerRequest) (entity.Customer, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(q string, page int, per_page int, entities string) ([]entity.Customer, error) {
	return s.repository.FindAll(q, page, per_page, entities)
}

func (s *service) Created(reqeuest CustomerRequest) (entity.Customer, error) {
	var customer entity.Customer
	customer = entity.Customer{
		Username: reqeuest.Username,
		Password: reqeuest.Password,
		Email:    reqeuest.Email,
		Contact:  reqeuest.Contact,
		Image:    reqeuest.Image,
		StoreId:  reqeuest.StoreId,
	}
	customers, err := s.repository.Created(customer)
	return customers, err
}
