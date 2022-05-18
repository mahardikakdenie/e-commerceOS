package product_customer

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(q string, entities string, page int, per_page int, category_id string, is_seller int, store_id string) ([]entity.Product, error)
	FindById(id uint, entities string, category_id string, store_id string) (entity.Product, error)
	FindByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(q string, entities string, page int, per_page int, category_id string, is_seller int, store_id string) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.
		Scopes(models.Search(q), models.Entities(entities), models.Paginate(page, per_page), models.ByCategoryId(category_id), models.Views(is_seller), models.ByStoreId(store_id)).
		Find(&products).Error
	return products, err
}

func (r *repository) FindById(id uint, entities string, category_id string, store_id string) (entity.Product, error) {
	var product entity.Product
	err := r.db.
		Scopes(models.Entities(entities), models.ByCategoryId(category_id), models.ByStoreId(store_id)).
		First(&product, id).Error
	return product, err
}

func (r *repository) FindByCategorySlug(slug string, entities string, page int, per_page int, is_seller int, store_id string) ([]entity.Product, error, error) {
	var products []entity.Product
	var category entity.Category
	category_err := r.db.Scopes(models.ByCategorySlug(slug)).First(&category).Error
	err := r.db.
		Scopes(models.Entities(entities), models.Paginate(page, per_page), models.SearchCategoryId(category.ID), models.Views(is_seller), models.ByStoreId(store_id)).
		Find(&products).Error
	return products, err, category_err
}
