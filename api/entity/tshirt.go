package entity

import "gorm.io/gorm"

type TShirt struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);unique_index"`
	Slug        string `gorm:"type:varchar(100);unique_index"`
	Description string `gorm:"type:varchar(355)"`
	Price       int    `gorm:"type:int"`
	Stock       int    `gorm:"type:int"`
	Image       string `gorm:"type:varchar(255)"`
	UserId      uint   `gorm:"foreignkey:UserID"`
	User        User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
