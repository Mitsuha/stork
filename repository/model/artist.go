package model

import (
	"time"
)

type Artist struct {
	ID        int        `json:"id" gorm:"column:id;primary_key"`
	Name      string     `json:"name" gorm:"column:name"`
	Image     string     `json:"image" gorm:"column:image"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
