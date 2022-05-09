package helper

import (
	"api/entity"
	"strconv"

	Slug "github.com/gosimple/slug"
)

func GeneratedProductStoreSlug(name string, data []entity.Product) string {
	var newSlugGenerated string
	if name != "" {
		newSlugGenerated = Slug.Make(name)
		for _, v := range data {
			if v.Slug == newSlugGenerated {
				var newData []entity.Product
				newData = append(newData, data...)
				is := strconv.Itoa(len(newData) + 1)
				newSlugGenerated = Slug.Make(name + "-" + is)
				return newSlugGenerated
			}
		}
	}

	return newSlugGenerated
}

func CheckProductStoreName(name string, data entity.Product) string {
	if name != "" {
		data.Name = name
	} else {
		name = data.Name
		data.Name = name
	}
	return name
}

func CheckProductStoreDescription(description string, data entity.Product) string {
	if description != "" {
		data.Description = description
	} else {
		description = data.Description
		data.Description = description
	}
	return description
}

func CheckProductStoreStock(stock int, data entity.Product) int {
	if stock != 0 {
		data.Stock = stock
	} else {
		stock = data.Stock
		data.Stock = stock
	}
	return stock
}

func CheckProductStorePrice(price float64, data entity.Product) float64 {
	if price != 0 {
		data.Price = price
	} else {
		price = data.Price
		data.Price = price
	}
	return price
}

func CheckProductStoreImage(image string, data entity.Product) string {
	if image != "" {
		data.Image = image
	} else {
		image = data.Image
		data.Image = image
	}
	return image
}

func CheckProductCategoryId(categoryId uint, data entity.Product) uint {
	if categoryId != 0 {
		data.CategoryId = categoryId
	} else {
		categoryId = data.CategoryId
		data.CategoryId = categoryId
	}

	return categoryId
}
