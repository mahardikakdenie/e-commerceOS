package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100);"`
	Image    string `gorm:"type:varchar(255);"`
	Contact  string `gorm:"type:varchar(255);"`
	StoreId  uint
	Store    Store
	Cart     []Cart
}
