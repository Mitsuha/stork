package model

import (
	"database/sql"
	"time"
)

type PersonalAccessToken struct {
	ID            int            `gorm:"column:id;primary_key"`
	TokenableType string         `gorm:"column:tokenable_type"`
	TokenableID   int            `gorm:"column:tokenable_id"`
	Name          string         `gorm:"column:name"`
	Token         string         `gorm:"column:token"`
	Abilities     sql.NullString `gorm:"column:abilities"`
	LastUsedAt    *time.Time     `gorm:"column:last_used_at"`
	CreatedAt     *time.Time     `gorm:"column:created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at"`
}
