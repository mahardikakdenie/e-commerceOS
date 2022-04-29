package store

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(q string, entities string) ([]entity.Store, error)
	Created(store entity.Store) (entity.Store, error)
	FindById(id int) (entity.Store, error)
	Updated(entity.Store) (entity.Store, error)
	Deleted(entity.Store) (entity.Store, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(q string, entities string) ([]entity.Store, error) {
	var stores []entity.Store
	if err := r.db.Scopes(models.Entities(entities), models.Search(q)).Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *repository) Created(store entity.Store) (entity.Store, error) {
	if err := r.db.Create(&store).Error; err != nil {
		return store, err
	}
	return store, nil
}

func (r *repository) FindById(id int) (entity.Store, error) {
	var store entity.Store
	if err := r.db.First(&store, id).Error; err != nil {
		return store, err
	}
	return store, nil
}

func (r *repository) Updated(store entity.Store) (entity.Store, error) {
	err := r.db.Save(&store).Error
	return store, err
}

func (r *repository) Deleted(store entity.Store) (entity.Store, error) {
	err := r.db.Delete(&store).Error
	return store, err
}
