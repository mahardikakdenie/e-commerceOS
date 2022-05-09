package helper

import (
	"api/entity"
	"strconv"

	Slug "github.com/gosimple/slug"
)

func GeneratedCategorySlug(name string, data []entity.Category) string {
	var newSlugGenerated string
	if name != "" {
		newSlugGenerated = Slug.Make(name)
		for _, v := range data {
			if v.Slug == newSlugGenerated {
				var newData []entity.Category
				newData = append(newData, data...)
				is := strconv.Itoa(len(newData) + 1)
				newSlugGenerated = Slug.Make(name + "-" + is)

				return newSlugGenerated
			}
		}
	}
	return newSlugGenerated
}

func CheckCategoryStoreName(name string, data entity.Category) string {
	if name != "" {
		data.Name = name
	} else {
		name = data.Name
		data.Name = name
	}

	return name
}

func CheckCategoryStoreDescription(description string, data entity.Category) string {
	if description != "" {
		data.Description = description
	} else {
		description = data.Description
		data.Description = description
	}
	return description
}

func CheckCategoryStoreSlug(slug string, data entity.Category) string {
	if slug != "" {
		data.Slug = slug
	} else {
		slug = data.Slug
		data.Slug = slug
	}
	return slug
}
