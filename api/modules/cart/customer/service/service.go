package service

import (
	"api/entity"
	"api/middleware"
	cart_customer "api/modules/cart/customer"
	"api/modules/customer"
	product "api/modules/product/store"
)

type service struct {
	repo         cart_customer.CartCussRepositoryInterface
	customerRepo customer.Repository
	productRepo  product.Repository
}

func NewService(repo cart_customer.CartCussRepositoryInterface, customerRepo customer.Repository, productRepo product.Repository) *service {
	return &service{
		repo:         repo,
		customerRepo: customerRepo,
		productRepo:  productRepo,
	}
}

func (s *service) IndexCartCustomer(entities string, customer_id uint, page int, per_page int) (carts []entity.Cart, err error, customer_err error) {
	data_customer, customer_err := s.repo.GetCustomerById(customer_id)
	if customer_err != nil {
		return nil, nil, customer_err
	}

	carts, err = s.repo.IndexCartCustomer(entities, data_customer.ID, data_customer.StoreId, page, per_page)
	if err != nil {
		return nil, err, nil
	}

	return carts, err, customer_err
}

func (s *service) InsertCart(cart entity.Cart, store_id uint, product_id uint) (entity.Cart, error) {
	// check customer is a valid customer
	if customer, err := s.customerRepo.Me(middleware.CustomerId); customer.StoreId != store_id {
		return entity.Cart{}, err
	}

	// check a product is a valid product
	if product, err := s.productRepo.GetId(product_id, store_id); product.StoreId != store_id {
		return entity.Cart{}, err
	}

	data, err := s.repo.InsertCart(cart)
	return data, err
}

func (s *service) DeleteCart(id uint, store_id uint, product_id uint) (entity.Cart, error) {
	// check custmer is a valid customer
	if customer, err := s.customerRepo.Me(middleware.CustomerId); customer.StoreId != store_id {
		return entity.Cart{}, err
	}

	// check product is a valid product
	if product, err := s.productRepo.GetId(product_id, store_id); product.StoreId != store_id {
		return entity.Cart{}, err
	}

	dataCart, err := s.repo.GetCartById(id)

	data, err := s.repo.DeleteCart(dataCart)
	return data, err
}
