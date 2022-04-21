package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	RoleId   uint `gorm:"foreignkey:RoleID"`
	Role     Role `json:"role"`
	Auth     []Auth
}
