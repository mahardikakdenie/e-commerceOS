package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);"`
	Slug        string    `gorm:"type:varchar(100);unique_index"`
	Description string    `gorm:"type:varchar(255)"`
	UserId      *uint     `gorm:"foreignkey:UserID"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product     []Product `json:"product"`
	StoreId     uint      `gorm:"foreignkey:StoreID"`
	Store       Store     `json:"store" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(100);"`
	Slug        string   `gorm:"type:varchar(100);unique_index"`
	Price       float64  `gorm:"type:float"`
	Image       string   `gorm:"type:varchar(255)"`
	Description string   `gorm:"type:varchar(455)"`
	Stock       int      `gorm:"type:int"`
	View        int      `gorm:"type:int;default:0"`
	CategoryId  uint     `gorm:"foreignkey:CategoryID"`
	UserId      *uint    `gorm:"foreignkey:UserID"`
	Category    Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User        User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StoreId     uint     `gorm:"foreignkey:StoreID"`
	Store       Store    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
