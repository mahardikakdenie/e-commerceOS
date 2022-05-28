package category_customer

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindCategoryByStore(store_id string, entities string) ([]entity.Category, error)
	FindCategoryBySlug(slug string, entities string) (entity.Category, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategoryByStore(store_id string, entities string) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Scopes(models.ByStoreId(store_id), models.Entities(entities)).Find(&categories).Error
	return categories, err
}

func (r *repository) FindCategoryBySlug(slug string, entities string) (entity.Category, error) {
	var category entity.Category
	err := r.db.Scopes(models.ByCategorySlug(slug), models.Entities(entities)).First(&category).Error
	return category, err
}
