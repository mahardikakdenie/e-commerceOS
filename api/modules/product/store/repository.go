package product_store

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(entities string, store_slug string) ([]entity.Product, error)
	Created(product entity.Product) (entity.Product, error)
	FindById(id uint, store_id string, entities string) (entity.Product, error)
	Updated(product entity.Product) (entity.Product, error)
	Deleted(product entity.Product) (entity.Product, error)
	GetId(id uint, store_id uint) (entity.Product, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(entities string, store_slug string) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Scopes(models.Entities(entities), models.ByStoreId(store_slug)).Find(&products).Error
	return products, err
}

func (r *repository) Created(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) FindById(id uint, store_id string, entities string) (entity.Product, error) {
	var product entity.Product
	err := r.db.Scopes(models.ByStoreId(store_id), models.Entities(entities)).First(&product, id).Error
	return product, err
}

func (r *repository) Updated(product entity.Product) (entity.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) Deleted(product entity.Product) (entity.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}

func (r *repository) GetId(id uint, store_id uint) (entity.Product, error) {
	var product entity.Product
	err := r.db.Where("store_id = ? ", store_id).First(&product, id).Error
	return product, err
}
