package models

import "gorm.io/gorm"

func HasCustommer(customerId uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if customerId != 0 {
			d.Preload("Customer", "id = ?", customerId)
		}
		return d
	}
}
