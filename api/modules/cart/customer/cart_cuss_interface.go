package cart_customer

import "api/entity"

type CartCussRepositoryInterface interface {
	IndexCartCustomer(entities string, CustomerId uint, store_id uint, page int, per_page int) (carts []entity.Cart, err error)
	GetCustomerById(customerId uint) (customer entity.Customer, err error)
	InsertCart(cart entity.Cart) (entity.Cart, error)
	DeleteCart(cart entity.Cart) (entity.Cart, error)
	GetCartById(id uint) (cart entity.Cart, err error)
}

type CartCussServiceInterface interface {
	IndexCartCustomer(entities string, customer_id uint, page int, per_page int) (carts []entity.Cart, err error, customer_err error)
	InsertCart(cart entity.Cart, store_id uint, product_id uint) (entity.Cart, error)
	DeleteCart(id uint, store_id uint, product_id uint) (entity.Cart, error)
}
