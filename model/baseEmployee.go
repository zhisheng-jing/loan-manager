package model

import (
	"time"
)

type BaseEmployee struct {
	ID  string  `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	IsActive    bool  `json:"is_actived"`
}
