package model

import (
	"database/sql"
	"time"
)

type Artist struct {
	ID        int            `gorm:"column:id;primary_key"`
	Name      string         `gorm:"column:name"`
	Image     sql.NullString `gorm:"column:image"`
	CreatedAt *time.Time     `gorm:"column:created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at"`
}
