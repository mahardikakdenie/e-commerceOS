package tshirt

import (
	"api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.TShirt, error)
	Created(tshirt entity.TShirt) (entity.TShirt, error)
	Updated(tshirt entity.TShirt) (entity.TShirt, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.TShirt, error) {
	var tshirts []entity.TShirt
	err := r.db.Preload("User").Find(&tshirts).Error
	return tshirts, err
}

func (r *repository) Created(tshirt entity.TShirt) (entity.TShirt, error) {
	err := r.db.Create(&tshirt).Error
	return tshirt, err
}

func (r *repository) Updated(tshirt entity.TShirt) (entity.TShirt, error) {
	err := r.db.Save(&tshirt).Error
	return tshirt, err
}
