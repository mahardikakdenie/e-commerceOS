package models

import "gorm.io/gorm"

func Search(search string) func(*gorm.DB) *gorm.DB {

	return func(d *gorm.DB) *gorm.DB {
		if search != "" {
			return d.Where("name LIKE ?", "%"+search+"%")
		}
		return d
	}
}
