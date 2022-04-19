package user

import (
	"api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Me(user_id int) (entity.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (entity.User, error) {
	var user entity.User
	err := r.db.Preload("Auth").Find(&user).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *repository) Me(user_id int) (entity.User, error) {
	var user entity.User
	err := r.db.Preload("Auth").Where("id = ?", user_id).Find(&user).Error
	return user, err
}
