package metadata

import (
	"github.com/mitsuha/stork/pkg/audio"
	"github.com/mitsuha/stork/pkg/lastfm"
	"github.com/mitsuha/stork/repository/model"
	"path/filepath"
	"strings"
)

type Service struct {
	lastfm *lastfm.Lastfm
}

func New(lastfm *lastfm.Lastfm) *Service {
	return &Service{lastfm: lastfm}
}

func (s *Service) Retrieve(m *model.Song, filename string, ref audio.Metadata) (*model.Artist, *model.Album) {
	if ref.HasTag() {
		fillModelWithMetadata(m, ref)
	} else if filename != "" {
		m.Title = strings.TrimSuffix(filename, filepath.Ext(filename))
		m.Length = ref.Duration()
	}

	artist, album := newArtistAndAlbum(ref)
	if s.lastfm != nil {
		_ = s.retrieveMoreMetadata(m, artist, album)
	}

	return artist, album
}

func fillModelWithMetadata(m *model.Song, metadata audio.Metadata) {
	if !metadata.HasTag() {
		return
	}

	m.Title, m.Length, m.Year, m.Genre, m.Track, m.Disc =
		metadata.Title(), metadata.Duration(), metadata.Year(), metadata.Genre(), metadata.Track(), metadata.Disc()
}

func newArtistAndAlbum(metadata audio.Metadata) (artist *model.Artist, album *model.Album) {
	artist, album = &model.Artist{Name: "Unknown artist"}, &model.Album{Name: "Unknown album"}

	if metadata.Artist() != "" {
		artist.Name = metadata.Artist()
	}

	if metadata.AlbumArtist() != "" {
		artist.Name = metadata.AlbumArtist()
	}

	if metadata.Album() != "" {
		album.Name = metadata.Album()
	}

	return
}

func (s *Service) retrieveMoreMetadata(m *model.Song, artist *model.Artist, album *model.Album) error {
	if artist.Name == "Unknown artist" {
		if tracks, err := s.lastfm.TrackSearch(m.Title, nil); err == nil && len(tracks) != 0 {
			m.Title, artist.Name = tracks[0].Name, tracks[0].Artist
		}
	}

	track, err := s.lastfm.TrackGetInfo(lastfm.Options{"track": {m.Title}, "artist": {artist.Name}})
	if err != nil {
		return err
	}

	album.Name, album.Cover = track.Album.Title, track.Album.GetImage()

	a, err := s.lastfm.ArtistGetInfo(lastfm.Options{"artist": {artist.Name}, "mbid": {track.Artist.Mbid}})
	if err != nil {
		return err
	}
	artist.Image = a.GetImage()

	return nil
}
