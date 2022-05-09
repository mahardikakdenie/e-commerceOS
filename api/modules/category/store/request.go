package category_store

type Request struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
	StoreId     uint   `json:"store_id"`
}
