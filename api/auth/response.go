package auth

type AuthResponse struct {
	ID        int    `json:"id"`
	AuthToken string `json:"token"`
	UserID    int    `json:"user_id"`
}
