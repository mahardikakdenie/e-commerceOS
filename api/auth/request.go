package auth

type AuthRequest struct {
	UserId    int    `json:"user_id" gorm:"foreignkey:UserID"`
	AuthToken string `json:"auth_token"`
}
