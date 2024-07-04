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

func (s *Songs) Create(song *model.Song, artist *model.Artist, album *model.Album) error {
	if artist != nil {
		artist, err := NewArtists(s.ctx).FindOrCreate(artist)
		if err != nil {
			return err
		}
		song.Artist, song.ArtistID = artist, artist.ID
	}

	if album != nil {
		album.ArtistID = song.ArtistID

		na, err := NewAlbums(s.ctx).FindOrCreate(album)
		if err != nil {
			return err
		}

		*album = *na

		song.Album, song.AlbumID = album, album.ID

		if artist != nil {
			song.Album.Artist = artist
		}
	}
	return dao.Song.WithContext(s.ctx).Create(song)
}
