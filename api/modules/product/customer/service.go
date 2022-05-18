package product_customer

import "api/entity"

type Service interface {
	FindAll(q string, entities string, page int, per_page int, category_id string, is_seller int, store_id string) ([]entity.Product, error)
	FindByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error, error)
	FindById(id uint, entities string, category_id string, store_id string) (entity.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(q string, entities string, page int, per_page int, category_id string, is_seller int, store_id string) ([]entity.Product, error) {
	return s.repository.
		FindAll(q, entities, page, per_page, category_id, is_seller, store_id)
}

func (s *service) FindByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error, error) {
	return s.repository.
		FindByCategorySlug(slug, entities, page, per_page, is_seller, store_id)
}

func (s *service) FindById(id uint, entities string, category_id string, store_id string) (entity.Product, error) {
	return s.repository.
		FindById(id, entities, category_id, store_id)
}
