package tshirt

import "gorm.io/gorm"

type tshirt struct {
	db *gorm.DB
}

type Thirt interface {
	FindAll()
}
