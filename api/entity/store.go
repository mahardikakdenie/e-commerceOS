package entity

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);unique_index"`
	Slug        string     `gorm:"type:varchar(100);unique_index"`
	Description string     `gorm:"type:varchar(355)"`
	Customer    []Customer `gorm:"customer_store;"`
	Cart        []Cart     `gorm:"cart_store;"`
}
