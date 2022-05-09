package product_store

type ProductStoreResponse struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  uint   `json:"category_id"`
	StoreId     uint   `json:"store_id"`
}
