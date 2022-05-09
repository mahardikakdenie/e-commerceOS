package customer_auth

import "api/entity"

type Service interface {
	FindAll() ([]entity.CustomerToken, error)
	GeneratedToken(request CustomerTokenRequest) (entity.CustomerToken, error)
	FindCustommer(username string) (entity.Customer, error)
	FindCustomerByToken(token string) (entity.CustomerToken, error)
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
