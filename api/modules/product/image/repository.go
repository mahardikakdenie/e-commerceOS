package image

import "api/entity"

type Repository interface {
	GetImageProduct(product_id uint) ([]entity.ProductImage, error)
	Insert(product entity.ProductImage) (entity.ProductImage, error)
}
