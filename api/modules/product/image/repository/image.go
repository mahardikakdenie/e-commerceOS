package repository

import (
	"api/entity"
	"api/models"

	// "api/vendor/gorm.io/gorm"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetImageProduct(product_id uint) ([]entity.ProductImage, error) {
	var product_image []entity.ProductImage
	err := r.db.Scopes(models.Entities("Product")).Where("product_id = ?", product_id).Find(&product_image).Error
	return product_image, err
}

func (r *repository) Insert(product entity.ProductImage) (entity.ProductImage, error) {
	err := r.db.Create(&product).Error
	return product, err
}
