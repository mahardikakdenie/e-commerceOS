package store

import (
	"time"

	"gorm.io/gorm"
)

type StoreResponses struct {
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Desc      string         `json:"description"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
