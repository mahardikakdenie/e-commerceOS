package auth

type AuthResponse struct {
	ID        int    `json:"id"`
	AuthToken string `json:"token"`
	UserID    uint   `json:"user_id"`
}
