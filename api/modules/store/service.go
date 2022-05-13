package store

import (
	"api/entity"
	h "api/helper"
)

type Service interface {
	FindAll(q string, entities string) ([]entity.Store, error)
	Created(store StoreRequest) (entity.Store, error)
	FindById(id int) (entity.Store, error)
	Update(request StoreRequest, id int) (entity.Store, error)
	Deleted(entity.Store) (entity.Store, error)
	FindBySlug(slug string) (entity.Store, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(q string, entities string) ([]entity.Store, error) {
	return s.repository.FindAll(q, entities)
}

func (s *service) Created(store StoreRequest) (entity.Store, error) {
	return s.repository.Created(entity.Store{
		Name:        store.Name,
		Slug:        store.Slug,
		Description: store.Description,
	})
}

func (s *service) FindById(id int) (entity.Store, error) {
	return s.repository.FindById(id)
}

func (s *service) Update(request StoreRequest, id int) (entity.Store, error) {
	store, err := s.repository.FindById(id)
	stores, err := s.repository.FindAll("", "")
	if err != nil {
	}

	store.Name = request.Name
	store.Slug = h.GeneratedStoreSlug(request.Name, stores)
	store.Description = request.Description

	updateData, err := s.repository.Updated(store)
	return updateData, err
}

func (s *service) Deleted(store entity.Store) (entity.Store, error) {
	return s.repository.Deleted(store)
}

func (s *service) FindBySlug(slug string) (entity.Store, error) {
	return s.repository.FindBySlug(slug)
}
