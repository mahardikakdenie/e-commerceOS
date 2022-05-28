package product_customer

import (
	"api/entity"
	"api/middleware"
	category_customer "api/modules/category/customer"
	"api/modules/customer"
	"errors"
	"strconv"
)

type Service interface {
	FindAll(q string, entities string, page int, per_page int, category_id string, is_seller int, store_id string) ([]entity.Product, error)
	FindByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error, error)
	FindById(id uint, entities string, category_id string, store_id string) (entity.Product, error)
	ShowProductStore(id uint, store_slug string) (entity.Product, error)
	FindProductByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error)
}

type service struct {
	repository   Repository
	customerRepo customer.Repository
	categoryRepo category_customer.Repository
}

func NewService(repository Repository, customerRepo customer.Repository, categoryRepo category_customer.Repository) *service {
	return &service{repository: repository, customerRepo: customerRepo, categoryRepo: categoryRepo}
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
	product, err := s.repository.
		FindById(id, entities, category_id, store_id)

	product.View += 1

	return product, err
}

func (s *service) ShowProductStore(id uint, store_slug string) (entity.Product, error) {

	customer, err := s.customerRepo.Me(middleware.CustomerId)
	if customer.Store.Slug != store_slug {
		return entity.Product{}, err
	}

	product, err := s.repository.FindById(id, "Category,User,Store", "", store_slug)

	if product.Store.Slug != store_slug {
		return entity.Product{}, err
	}

	return product, err
}

func (s *service) FindProductByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error) {

	// check Customer is valid
	customer, err := s.customerRepo.Me(middleware.CustomerId)
	storeId, _ := strconv.Atoi(store_id)
	if customer.StoreId != uint(storeId) {
		return []entity.Product{}, errors.New("Customer is not valid")
	}

	// check category is valid
	category, err := s.categoryRepo.FindCategoryBySlug(slug, entities)
	if category.StoreId != uint(storeId) {
		return []entity.Product{}, errors.New("Category is not valid")
	}

	// execute code
	category_id := strconv.FormatUint(uint64(category.ID), 10)
	products, err := s.repository.FindAll(slug, entities, page, per_page, category_id, is_seller, store_id)

	return products, err
}
