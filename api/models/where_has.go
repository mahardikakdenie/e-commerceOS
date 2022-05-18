package models

import "gorm.io/gorm"

func ByCustomer(customerId uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if customerId != 0 {
			d.Preload("Customer", "id = ?", customerId)
		}
		return d
	}
}

func ByStore(store_id uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if store_id != 0 {
			d.Preload("Store", "id = ?", store_id)
		}
		return d
	}
}

func ByStoreId(store_id string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if store_id != "" {
			d.Where("store_id = ?", store_id)
		}
		return d
	}
}

func ByCategoryId(category_id string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if category_id != "" {
			d.Where("category_id = ?", category_id)
		}
		return d
	}
}

func SearchCategoryId(category_id uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if category_id != 0 {
			d.Where("category_id = ?", category_id)
		}
		return d
	}
}

func ByCategorySlug(category_slug string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if category_slug != "" {
			d.Where("slug = ?", category_slug)
		}
		return d
	}
}

func CartByCustomerId(customerId uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if customerId != 0 {
			d.Preload("Customer", "id = ?", customerId)
		}
		return d
	}
}
