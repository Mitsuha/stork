package model

type PlaylistSong struct {
	ID         int    `json:"id" gorm:"primarykey"`
	PlaylistID int    `json:"playlist_id" gorm:"column:playlist_id"`
	SongID     string `json:"song_id" gorm:"column:song_id"`
}

func (p *PlaylistSong) TableName() string {
	return "playlist_song"
}
