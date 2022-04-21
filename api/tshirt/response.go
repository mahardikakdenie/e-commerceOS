package tshirt

import "api/entity"

type TshirtResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Description string      `json:"description"`
	Price       int         `json:"price"`
	Stock       int         `json:"stock"`
	UserId      uint        `json:"author_id"`
	User        entity.User `json:"author"`
	Link        Link        `json:"link"`
}

type Link struct {
	Image string `json:"url_image"`
}
