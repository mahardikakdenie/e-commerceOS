package cart_customer

import (
	"api/entity"
	"api/middleware"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(q string, entities string, page int, per_page int) (carts []entity.Cart, err error, customer_err error)
	GetDataCartByCustomerId(customerId uint) (carts []entity.Cart, err error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(q string, entities string, page int, per_page int) (carts []entity.Cart, err error, customer_err error) {
	var customer entity.Customer
	customer_err = r.db.
		Find(&customer, middleware.CustomerId).
		Error
	err = r.db.
		Scopes(models.Entities(entities), models.ByStore(customer.StoreId), models.CartByCustomerId(customer.ID)).
		Find(&carts).Error
	return carts, err, customer_err
}

func (r *repository) GetDataCartByCustomerId(customerId uint) (carts []entity.Cart, err error) {
	err = r.db.
		Scopes(models.CartByCustomerId(customerId)).
		Find(&carts).Error
	return carts, err
}
