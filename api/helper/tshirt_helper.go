package helper

import (
	"api/entity"
	"strconv"

	Slug "github.com/gosimple/slug"
)

func GeneratedSlug(name string, data []entity.TShirt) string {
	var newSlugGenerated string
	if name != "" {
		newSlugGenerated = Slug.Make(name)
		for _, v := range data {
			if v.Slug == newSlugGenerated {
				var newData []entity.TShirt
				newData = append(newData, data...)
				is := strconv.Itoa(len(newData) + 1)
				newSlugGenerated = Slug.Make(name + "-" + is)

				return newSlugGenerated
			}
		}
	}
	return newSlugGenerated
}

func CheckerName(name string, data entity.TShirt) string {
	if name != "" {
		data.Name = name
	} else {
		name = data.Name
		data.Name = name
	}
	return name
}

func CheckerPrice(price int, data entity.TShirt) int {
	if price != 0 {
		data.Price = price
	} else {
		price = data.Price
		data.Price = price
	}
	return price
}

func CheckStock(stock int, data entity.TShirt) int {
	if stock != 0 {
		data.Stock = stock
	} else {
		stock = data.Stock
		data.Stock = stock
	}
	return stock
}

func CheckDescription(description string, data entity.TShirt) string {
	if description != "" {
		data.Description = description
	} else {
		description = data.Description
		data.Description = description
	}
	return description
}

func CheckImage(image string, data entity.TShirt) string {
	if image != "" {
		data.Image = image
	} else {
		image = data.Image
		data.Image = image
	}
	return image
}

func CheckerSlug(slug string, data entity.TShirt) string {
	if slug != "" {
		data.Slug = slug
	} else {
		slug = data.Slug
		data.Slug = slug
	}
	return slug
}
