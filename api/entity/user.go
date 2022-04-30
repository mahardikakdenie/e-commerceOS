package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Image    string
	Contact  string
	RoleId   uint `gorm:"foreignkey:RoleID"`
	Role     Role `json:"role"`
	Auth     []Auth
	StoreId  *uint `gorm:"foreignkey:StoreID:constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Store    Store `grom:"foreignKey:StoreID:constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
