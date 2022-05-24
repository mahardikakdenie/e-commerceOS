package customer

import (
	"api/entity"
	"api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(q string, page int, per_page int, entites string) ([]entity.Customer, error)
	FindById(id int, entities string) (entity.Customer, error)
	Created(customer entity.Customer) (entity.Customer, error)
	Updated(customer entity.Customer) (entity.Customer, error)
	Deleted(customer entity.Customer) (entity.Customer, error)
	Me(id uint) (entity.Customer, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(q string, page int, per_page int, entites string) ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Scopes(models.Paginate(page, per_page), models.SearchCustomerUsername(q), models.Entities(entites)).Find(&customers).Error
	return customers, err
}

func (r *repository) FindById(id int, entities string) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Scopes(models.Entities(entities)).First(&customer, id).Error
	return customer, err
}

func (r *repository) Me(id uint) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("id = ? ", id).Preload("Store").First(&customer).Error
	return customer, err
}

func (r *repository) Created(customer entity.Customer) (entity.Customer, error) {
	err := r.db.Create(&customer).Error
	return customer, err
}

func (r *repository) Updated(customer entity.Customer) (entity.Customer, error) {
	err := r.db.Save(&customer).Error
	return customer, err
}

func (r *repository) Deleted(customer entity.Customer) (entity.Customer, error) {
	err := r.db.Delete(&customer).Error
	return customer, err
}
