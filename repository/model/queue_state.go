package model

import (
	"github.com/mitsuha/stork/repository/model/types"
	"time"
)

type QueueState struct {
	ID               int               `json:"id" gorm:"column:id;primary_key"`
	UserID           int               `json:"user_id" gorm:"column:user_id"`
	SongIds          types.StringArray `json:"song_ids" gorm:"column:song_ids;type:text"`
	CurrentSongID    string            `json:"current_song_id" gorm:"column:current_song_id"`
	PlaybackPosition int               `json:"playback_position" gorm:"column:playback_position"`
	CreatedAt        time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time         `json:"updated_at" gorm:"column:updated_at"`
}
