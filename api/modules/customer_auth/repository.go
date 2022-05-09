package customer_auth

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.CustomerToken, error)
	GeneratedToken(entity.CustomerToken) (entity.CustomerToken, error)
	FindCustomer(username string) (entity.Customer, error)
	FindCustomerByToken(token string) (entity.CustomerToken, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.CustomerToken, error) {
	var customerTokens []entity.CustomerToken
	err := r.db.Scopes(models.Entities("Customer")).Find(&customerTokens).Error
	return customerTokens, err
}

func (r *repository) GeneratedToken(customerToken entity.CustomerToken) (entity.CustomerToken, error) {
	err := r.db.Create(&customerToken).Error
	return customerToken, err
}

func (r *repository) FindCustomer(username string) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Scopes(models.SearchCustomerUsername(username)).First(&customer).Error
	return customer, err
}

func (r *repository) FindCustomerByToken(token string) (entity.CustomerToken, error) {
	var customer entity.CustomerToken
	err := r.db.Scopes(models.SearchCustomerToken(token), models.Entities("Customer")).First(&customer).Error
	return customer, err
}
