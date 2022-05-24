package repository

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return *&repository{
		db: db,
	}
}

func (r *repository) IndexCartCustomer(entities string, CustomerId uint, store_id uint, page int, per_page int) (carts []entity.Cart, err error) {
	err = r.db.
		Scopes(models.Entities(entities),
			models.WhereCustomerId(CustomerId),
			models.Paginate(page, per_page),
			models.WhereStoreId(store_id),
		).
		Find(&carts).Error
	return carts, err
}

func (r *repository) GetCustomerById(customerId uint) (customer entity.Customer, err error) {
	err = r.db.
		Find(&customer, customerId).Error
	return customer, err
}

func (r *repository) InsertCart(cart entity.Cart) (entity.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) DeleteCart(data entity.Cart) (entity.Cart, error) {
	var cart entity.Cart
	err := r.db.Delete(&data).Error

	return cart, err
}

func (r *repository) GetCartById(id uint) (entity.Cart, error) {
	var cart entity.Cart
	err := r.db.
		Find(&cart, id).Error
	return cart, err
}
