package order

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(q string, page int, par_page int, rntities string) ([]entity.Order, error)
	FindById(id int) (entity.Order, error)
	Created(order entity.Order) (entity.Order, error)
	Updated(order entity.Order) (entity.Order, error)
	Deleted(order entity.Order) (entity.Order, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(q string, page int, par_page int, entities string) ([]entity.Order, error) {
	var order []entity.Order
	err := r.db.Scopes(models.Paginate(page, par_page), models.Search(q), models.Entities(entities)).Find(&order).Error
	return order, err
}

func (r *repository) FindById(id int) (entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("User").Preload("TShirt").First(&order, id).Error
	return order, err
}

func (r *repository) Created(order entity.Order) (entity.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *repository) Updated(order entity.Order) (entity.Order, error) {
	err := r.db.Preload("User").Preload("TShirt").Save(&order).Error
	return order, err
}

func (r *repository) Deleted(order entity.Order) (entity.Order, error) {
	err := r.db.Delete(&order).Error
	return order, err
}
