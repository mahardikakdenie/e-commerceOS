package category_store

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(entities string, store_id uint) ([]entity.Category, error)
	Created(category entity.Category) (entity.Category, error)
	Show(id uint, entities string) (entity.Category, error)
	Updated(category entity.Category) (entity.Category, error)
	Deleted(category entity.Category) (entity.Category, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(entities string, store_id uint) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Scopes(models.Entities(entities), models.ByStore(store_id)).Find(&categories).Error
	return categories, err
}

func (r *repository) Created(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *repository) Show(id uint, entities string) (entity.Category, error) {
	var category entity.Category
	err := r.db.Scopes(models.Entities(entities)).Find(&category, id).Error
	return category, err
}

func (r *repository) Updated(category entity.Category) (entity.Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *repository) Deleted(category entity.Category) (entity.Category, error) {
	err := r.db.Delete(&category).Error
	return category, err
}
