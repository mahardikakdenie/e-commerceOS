package auth

import "api/entity"

type AuthResponse struct {
	ID        uint        `json:"id"`
	AuthToken string      `json:"token"`
	UserID    uint        `json:"user_id"`
	User      entity.User `json:"user"`
}
