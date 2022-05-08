package entity

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UserId    uint   `json:"user_id" gorm:"foreignkey:UserID"`
	User      User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthToken string `json:"auth_token"`
}
