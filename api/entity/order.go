package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TShirtID   uint     `gorm:"foreignkey:TShirtID"`
	TShirt     TShirt   `json:"tshirt" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity   int      `gorm:"type:int"`
	Status     string   `gorm:"type:varchar(100)"`
	Address    string   `gorm:"type:varchar(255)"`
	CustomerID *uint    `gorm:"foreignkey:CustomerID"`
	UserID     *uint    `gorm:"foreignkey:UserID"`
	Customer   Customer `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User       User
}
