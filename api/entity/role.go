package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	User []User
}
