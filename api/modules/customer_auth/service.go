package customer_auth

import (
	"api/entity"
)

type Service interface {
	FindAll() ([]entity.CustomerToken, error)
	GeneratedToken(request CustomerTokenRequest) (entity.CustomerToken, error)
	FindCustommer(username string) (entity.Customer, error)
	FindCustomerByToken(token string) (entity.CustomerToken, error)
	Logout(token string) (entity.CustomerToken, error)
	FindStoreById(id uint) (entity.Store, error)
	FindStoreBySlug(slug string) (entity.Store, error)
	Register(customer CustomerRequest) (entity.Customer, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.CustomerToken, error) {
	return s.repository.FindAll()
}

func (s *service) GeneratedToken(request CustomerTokenRequest) (entity.CustomerToken, error) {
	var customerToken entity.CustomerToken
	customerToken = entity.CustomerToken{
		CustomerId: uint(request.CustomerId),
		AuthToken:  request.AuthToken,
	}
	customerTokens, err := s.repository.GeneratedToken(customerToken)
	return customerTokens, err
}

func (s *service) FindCustommer(username string) (entity.Customer, error) {
	return s.repository.FindCustomer(username)
}

func (s *service) FindCustomerByToken(token string) (entity.CustomerToken, error) {
	return s.repository.FindCustomerByToken(token)
}

func (s *service) Logout(token string) (entity.CustomerToken, error) {
	tokens, err := s.repository.FindByToken(token)
	if err != nil {
		return tokens, err
	}
	data, err := s.repository.Logout(tokens)
	return data, err
}

func (s *service) FindStoreById(id uint) (entity.Store, error) {
	return s.repository.FindStoreById(id)
}

func (s *service) FindStoreBySlug(slug string) (entity.Store, error) {
	return s.repository.FindStoreBySlug(slug)
}

func (s *service) Register(customer CustomerRequest) (entity.Customer, error) {
	var customerEntity entity.Customer
	customerEntity = entity.Customer{
		Username: customer.Username,
		Password: customer.Password,
		Email:    customer.Email,
		Contact:  customer.Contact,
		StoreId:  customer.StoreId,
	}
	return s.repository.Register(customerEntity)
}
