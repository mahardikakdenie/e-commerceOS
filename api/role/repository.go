package role

import (
	"api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Role, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Find(&roles).Error
	return roles, err
}
