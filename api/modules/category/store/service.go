package category_store

import (
	"api/entity"
	"api/helper"
	"api/middleware"
)

type Service interface {
	FindAll(entities string, store_id uint) ([]entity.Category, error)
	Created(Request Request) (entity.Category, error)
	Show(id uint, entities string) (entity.Category, error)
	Updated(id uint, request Request) (entity.Category, error)
	Deleted(id uint) (entity.Category, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(entities string, store_id uint) ([]entity.Category, error) {
	return s.repository.FindAll(entities, store_id)
}

func (s *service) Created(request Request) (entity.Category, error) {
	var category entity.Category
	category = entity.Category{
		StoreId:     *middleware.StoreId,
		Name:        request.Name,
		Slug:        request.Slug,
		Description: request.Description,
		UserId:      &middleware.UserId,
	}

	return s.repository.Created(category)
}

func (s *service) Show(id uint, entities string) (entity.Category, error) {
	return s.repository.Show(id, entities)
}

func (s *service) Updated(id uint, request Request) (entity.Category, error) {
	category, err := s.repository.Show(id, "")
	if err != nil {
		return category, err
	}
	category.Name = helper.CheckCategoryStoreName(request.Name, category)
	category.Slug = request.Slug
	category.Description = helper.CheckCategoryStoreDescription(request.Description, category)
	category.UserId = &middleware.UserId
	category.StoreId = *middleware.StoreId
	return s.repository.Updated(category)
}

func (s *service) Deleted(id uint) (entity.Category, error) {
	category, _ := s.repository.Show(id, "")
	return s.repository.Deleted(category)
}
