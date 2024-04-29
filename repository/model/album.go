package model

import "time"

type Album struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	ArtistID  int        `json:"artist_id" gorm:"column:artist_id"`
	Name      string     `json:"name" gorm:"type:mediumtext;column:name"`
	Cover     string     `json:"cover" gorm:"type:varchar(191);default:'';column:cover"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	Artist    Artist     `json:"artist" gorm:"foreignKey:artist_id"`
}
