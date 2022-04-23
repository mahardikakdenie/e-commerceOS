package cart

import (
	"api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Cart, error)
	Created(entity.Cart) (entity.Cart, error)
	Updated(entity.Cart) (entity.Cart, error)
	FindById(id int) (entity.Cart, error)
	Deleted(entity.Cart) (entity.Cart, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Cart, error) {
	var cart []entity.Cart
	err := r.db.Find(&cart).Error
	return cart, err
}

func (r *repository) Created(cart entity.Cart) (entity.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) Updated(cart entity.Cart) (entity.Cart, error) {
	err := r.db.Save(&cart).Error
	return cart, err
}

func (r *repository) Deleted(cart entity.Cart) (entity.Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}

func (r *repository) FindById(id int) (entity.Cart, error) {
	var cart entity.Cart
	err := r.db.Find(&cart, id).Error
	return cart, err
}
