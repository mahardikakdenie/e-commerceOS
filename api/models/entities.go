package models

import (
	"strings"

	"gorm.io/gorm"
)

func Entities(entities string) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		if entities != "" {
			new_array := strings.Split(entities, ",")
			for _, v := range new_array {
				if v != " " {
					db.Preload(v)
				}
			}
		}

		return db
	}
}
