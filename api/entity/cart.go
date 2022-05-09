package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   uint   `gorm:"foreignkey:UserID"`
	User     User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity int    `gorm:"type:int"`
	Status   string `gorm:"type:varchar(100)"`
}
