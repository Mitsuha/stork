package model

import (
	"time"
)

type Playlist struct {
	ID        int       `json:"id" gorm:"primarykey"`
	UserID    int       `json:"user_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"type:mediumtext;not null"`
	Rules     string    `json:"rules" gorm:"type:text"`
	FolderID  *string   `json:"folder_id" gorm:"type:varchar(36)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
