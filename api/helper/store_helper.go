package helper

import (
	"api/entity"
	"strconv"

	Slug "github.com/gosimple/slug"
)

func GeneratedStoreSlug(name string, data []entity.Store) string {
	var slug string
	if name != "" {
		slug = Slug.Make(name)
		for _, v := range data {
			if v.Slug == slug {
				var newData []entity.Store
				newData = append(newData, data...)
				is := strconv.Itoa(len(newData) + 1)
				slug = Slug.Make(name + "-" + is)
				return slug
			}
		}
	}

	return slug
}

func CheckStoreName(name string, data entity.Store) string {
	if name != "" {
		data.Name = name
	} else {
		name = data.Name
		data.Name = name
	}
	return name
}

func CheckStoreDescription(description string, data entity.Store) string {
	if description != "" {
		data.Description = description
	} else {
		description = data.Description
		data.Description = description
	}
	return description
}

func CheckerStroreSlug(slug string, data entity.Store) string {
	if slug != "" {
		data.Slug = slug
	} else {
		slug = data.Slug
		data.Slug = slug
	}
	return slug
}
