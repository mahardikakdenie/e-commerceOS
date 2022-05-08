package tshirt

type TshirtRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Image       string `json:"url_image"`
	UserId      uint   `json:"user_id"`
}
