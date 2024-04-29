package model

import "time"

type Interaction struct {
	ID           int64     `gorm:"column:id;primary_key"`
	UserID       int       `gorm:"column:user_id"`
	SongID       string    `gorm:"column:song_id"`
	Liked        bool      `gorm:"column:liked"`
	PlayCount    int       `gorm:"column:play_count"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	LastPlayedAt time.Time `gorm:"column:last_played_at"`
}
