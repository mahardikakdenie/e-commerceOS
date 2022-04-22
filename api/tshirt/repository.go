package tshirt

import (
	"api/entity"
	models "api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(q string, page int, page_size int) ([]entity.TShirt, error)
	FindById(id int) (entity.TShirt, error)
	Created(tshirt entity.TShirt) (entity.TShirt, error)
	Updated(tshirt entity.TShirt) (entity.TShirt, error)
	Delete(tshirt entity.TShirt) (entity.TShirt, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(q string, page int, page_size int) ([]entity.TShirt, error) {
	var tshirts []entity.TShirt
	// rQuery := http.Request{}
	err := r.db.Scopes(models.Search(q), models.Paginate(page, page_size)).Preload("User").Find(&tshirts).Error
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

func (r *repository) FindById(id int) (entity.TShirt, error) {
	var tshirt entity.TShirt
	err := r.db.Preload("User").First(&tshirt, id).Error
	return tshirt, err
}

func (r *repository) Delete(tshirt entity.TShirt) (entity.TShirt, error) {
	err := r.db.Delete(&tshirt).Error
	return tshirt, err
}
