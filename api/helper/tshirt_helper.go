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
