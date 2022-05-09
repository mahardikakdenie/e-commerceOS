package product_store

import (
	"api/entity"
	"api/helper"
	"api/middleware"
)

type Service interface {
	FindAll(entities string) ([]entity.Product, error)
	Created(request ProductStoreRequest) (entity.Product, error)
	Updated(request ProductStoreRequest, id uint) (entity.Product, error)
	FindById(id uint) (entity.Product, error)
	Deleted(id uint) (entity.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(entities string) ([]entity.Product, error) {
	return s.repository.FindAll(entities)
}

func (s *service) FindById(id uint) (entity.Product, error) {
	return s.repository.FindById(id)
}

func (s *service) Created(request ProductStoreRequest) (entity.Product, error) {
	var product entity.Product
	products, err := s.repository.FindAll("")
	product = entity.Product{
		Name:        request.Name,
		Slug:        helper.GeneratedProductStoreSlug(request.Name, products),
		Description: request.Description,
		Stock:       request.Stock,
		Image:       request.Image,
		Price:       float64(request.Price),
		StoreId:     *middleware.StoreId,
		UserId:      &middleware.UserId,
		CategoryId:  request.CategoryId,
	}
	data, err := s.repository.Created(product)
	return data, err
}

func (s *service) Updated(request ProductStoreRequest, id uint) (entity.Product, error) {
	product, err := s.repository.FindById(id)
	products, err := s.repository.FindAll("")

	product.Name = helper.CheckProductStoreName(request.Name, product)
	product.Slug = helper.GeneratedProductStoreSlug(request.Name, products)
	product.Description = helper.CheckProductStoreDescription(request.Description, product)
	product.Stock = helper.CheckProductStoreStock(request.Stock, product)
	product.Image = helper.CheckProductStoreImage(request.Image, product)
	product.Price = helper.CheckProductStorePrice(float64(request.Price), product)
	product.StoreId = request.StoreId
	product.UserId = &request.UserId
	product.CategoryId = helper.CheckProductCategoryId(request.CategoryId, product)

	dataUpdate, err := s.repository.Updated(product)
	return dataUpdate, err
}

func (s *service) Deleted(id uint) (entity.Product, error) {
	product, _ := s.repository.FindById(id)
	return s.repository.Deleted(product)
}
