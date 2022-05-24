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

func SearchCustomerUsername(search string) func(*gorm.DB) *gorm.DB {

	return func(d *gorm.DB) *gorm.DB {
		if search != "" {
			return d.Where("username = ?", search).Or("email = ?", search)
		}
		return d
	}
}

func WhereCustommer(username string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if username != "" {
			return d.Where("username = ?", username).Or("email = ?", username)
		}
		return d
	}
}

func SearchCustomerToken(search string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if search != "" {
			return d.Where("auth_token LIKE ?", "%"+search+"%")
		}
		return d
	}
}

func SearchStoreSlug(slug string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if slug != "" {
			return d.Where("slug LIKE ?", "%"+slug+"%")
		}
		return d
	}
}

func Views(views int) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if views == 1 {
			return d.Order("view desc")
		}
		return d
	}
}

func WhereCustomerId(customer_id uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if customer_id != 0 {
			return d.Where("customer_id = ?", customer_id)
		}
		return d
	}
}

func WhereStoreId(storeId uint) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if storeId != 0 {
			return d.Where("store_id = ?", storeId)
		}
		return d
	}
}
