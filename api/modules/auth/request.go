package auth

type AuthRequest struct {
	UserId    uint   `json:"user_id" gorm:"foreignkey:UserID"`
	AuthToken string `json:"auth_token"`
}
