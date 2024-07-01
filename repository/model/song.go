package model

import (
	"time"
)

type SongFrom int8

const (
	SFromLocal SongFrom = iota
	SFromCloud
)

type Song struct {
	ID          string       `json:"id" gorm:"column:id;primary_key"`
	AlbumID     int          `json:"album_id" gorm:"column:album_id"`
	Title       string       `json:"title" gorm:"column:title"`
	Length      float64      `json:"length" gorm:"column:length"`
	Track       int          `json:"track" gorm:"column:track"`
	Disc        int          `json:"disc" gorm:"column:disc"`
	Lyrics      string       `json:"lyrics" gorm:"column:lyrics"`
	Path        string       `json:"path" gorm:"column:path"`
	Mtime       int          `json:"mtime" gorm:"column:mtime"`
	CreatedAt   time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"column:updated_at"`
	ArtistID    int          `json:"artist_id" gorm:"column:artist_id"`
	Year        int          `json:"year" gorm:"column:year"`
	Genre       string       `json:"genre" gorm:"column:genre"`
	Album       *Album       `json:"album" gorm:"foreignKey:album_id"`
	Artist      *Artist      `json:"artist" gorm:"foreignKey:artist_id"`
	Interaction *Interaction `json:"interaction" gorm:"foreignKey:song_id"`
	From        SongFrom     `json:"from" gorm:"column:from"`
}

type CountAndLength struct {
	Count  int     `json:"count"`
	Length float64 `json:"length"`
}
