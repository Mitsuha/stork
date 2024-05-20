package repository

import (
	"context"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Songs struct {
	ctx context.Context
}

func NewSongs(ctx context.Context) *Songs {
	return &Songs{ctx: ctx}
}

func (s *Songs) Create(song *model.Song) error {
	if song.Artist != nil {
		artist, err := NewArtists(s.ctx).FindOrCreate(song.Artist)
		if err != nil {
			return err
		}
		song.Artist, song.ArtistID = artist, artist.ID
	}

	if song.Album != nil {
		song.Album.ArtistID = song.ArtistID

		album, err := NewAlbums(s.ctx).FindOrCreate(song.Album)
		if err != nil {
			return err
		}
		song.Album, song.AlbumID, song.Album.Artist = album, album.ID, song.Artist
	}
	return dao.Song.WithContext(s.ctx).Create(song)
}
