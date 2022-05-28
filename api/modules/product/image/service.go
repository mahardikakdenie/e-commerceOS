package image

import (
	"api/entity"
	rr "api/modules/product/image/RR"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetImageProduct(ctx *gin.Context, product_id uint) ([]entity.ProductImage, error)
	InsertImage(ctx *gin.Context, request rr.ProductImageRequest) (entity.ProductImage, error)
}
