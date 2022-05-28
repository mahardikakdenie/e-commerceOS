package user

import (
	"api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Me(user_id uint) (entity.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var user []entity.User                                        // Fetching data from FOlder Entity User
	err := r.db.Preload("Auth").Preload("Role").Find(&user).Error // quuery for fecth data from database
	return user, err
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *repository) Me(user_id uint) (entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role,Store").Where("id = ?", user_id).Find(&user).Error
	return user, err
}
