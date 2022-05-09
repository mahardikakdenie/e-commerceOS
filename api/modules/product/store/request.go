package product_store

type ProductStoreRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
	CategoryId  uint   `json:"category_id"`
	StoreId     uint   `json:"store_id"`
	UserId      uint   `json:"user_id"`
}
