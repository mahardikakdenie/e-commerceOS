package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	// UserID     uint     `gorm:"foreignkey:UserID"`
	// User       User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity   int      `gorm:"type:int"`
	Status     string   `gorm:"type:varchar(100)"`
	StoreID    uint     `gorm:"foreignkey:StoreID"`
	Store      Store    `json:"store" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CustomerId uint     `gorm:"foreignkey:CustomerId"`
	Customer   Customer `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
