package service

import (
	"api/entity"
	"api/helper"
	"api/middleware"
	product_customer "api/modules/product/customer"
	"api/modules/product/image"
	rr "api/modules/product/image/RR"
	"api/modules/user"
	"errors"

	"github.com/gin-gonic/gin"
)

type service struct {
	repo        image.Repository
	productRepo product_customer.Repository
	userRepo    user.Repository
}

func NewService(repo image.Repository, productRepo product_customer.Repository, userRepo user.Repository) *service {
	return &service{
		repo:        repo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (s *service) GetImageProduct(ctx *gin.Context, product_id uint) ([]entity.ProductImage, error) {

	// check a valid Image
	product, err := s.productRepo.FindById(product_id, "Product", "", "")
	if product.ID == 0 {
		helper.Exception(ctx, false, "Product not found", err)
	}

	product_image, err := s.repo.GetImageProduct(product.ID)

	return product_image, err
}

func (s *service) InsertImage(ctx *gin.Context, request rr.ProductImageRequest) (entity.ProductImage, error) {

	// check a valid Image
	product, err := s.productRepo.FindById(request.ProductID, "", "", "")
	if product.ID == 0 {
		helper.Exception(ctx, false, "Product not found", err)
		return entity.ProductImage{}, errors.New("Product not found")
	}

	user, err := s.userRepo.Me(middleware.UserId)
	if *user.StoreId != request.StoreId {
		return entity.ProductImage{}, errors.New("You are not owner")
	}

	data := entity.ProductImage{
		ProductId: request.ProductID,
		Image:     request.Image,
		StoreId:   request.StoreId,
		UserId:    &middleware.UserId,
	}

	product_image, err := s.repo.Insert(data)

	return product_image, err
}
