package category_customer

import "api/entity"

type Service interface {
	FindCategoryByStore(store_id string, entities string) ([]entity.Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCategoryByStore(store_id string, entities string) ([]entity.Category, error) {
	return s.repository.FindCategoryByStore(store_id, entities)
}
