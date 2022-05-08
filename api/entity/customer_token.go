package entity

import "gorm.io/gorm"

type CustomerToken struct {
	gorm.Model
	AuthToken  string `gorm:"type:varchar(100);unique_index"`
	CustomerId uint   `json:"customer_id" gorm:"foreignkey:CustomerID"`
	Customer   Customer
}
