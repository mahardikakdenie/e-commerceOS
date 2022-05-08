package customer

import (
	"time"

	"gorm.io/gorm"
)

type CustomerResponse struct {
	Id        uint           `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Contact   string         `json:"contact"`
	Image     string         `json:"image"`
	UpdateAt  time.Time      `json:"update_at"`
	CreateAt  time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
