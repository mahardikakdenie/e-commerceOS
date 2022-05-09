package auth

import (
	"api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Login(auth entity.Auth) (entity.Auth, error)
	Register(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByToken(token string) (entity.Auth, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(auth entity.Auth) (entity.Auth, error) {
	// var user entity.Auth
	err := r.db.Create(&auth).Error
	return auth, err
}

func (r *repository) Register(user entity.User) (entity.User, error) {
	// var user entity.User
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *repository) FindByToken(token string) (entity.Auth, error) {
	var auth entity.Auth
	err := r.db.Preload("User").Where("auth_token = ?", token).Find(&auth).Error
	return auth, err
}
