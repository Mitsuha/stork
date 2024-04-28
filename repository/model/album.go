package model

import "time"

type Album struct {
	ID        int
	ArtistID  int
	Name      string
	Cover     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
