package model

import (
	"github.com/mitsuha/stork/repository/model/types"
	"time"
)

type PersonalAccessToken struct {
	ID            int               `gorm:"column:id;primary_key"`
	TokenableType string            `gorm:"column:tokenable_type"`
	TokenableID   int               `gorm:"column:tokenable_id"`
	Name          string            `gorm:"column:name"`
	Token         string            `gorm:"column:token"`
	Abilities     types.StringArray `gorm:"column:abilities;type:varchar(255)"`
	LastUsedAt    *time.Time        `gorm:"column:last_used_at"`
	CreatedAt     *time.Time        `gorm:"column:created_at"`
	UpdatedAt     *time.Time        `gorm:"column:updated_at"`
}
