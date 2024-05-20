package overview

import (
	"github.com/mitsuha/stork/repository"
	"github.com/mitsuha/stork/repository/model"
)

type DataResponse struct {
	Settings            any             `json:"settings"`
	Playlists           string          `json:"playlists"`
	PlaylistFolders     string          `json:"playlist_folders"`
	CurrentUser         *UserResp       `json:"current_user"`
	UseLastFm           bool            `json:"use_last_fm"`
	UseSpotify          bool            `json:"use_spotify"`
	UseYouTube          bool            `json:"use_you_tube"`
	UseITunes           bool            `json:"use_i_tunes"`
	AllowDownload       bool            `json:"allow_download"`
	SupportsTranscoding bool            `json:"supports_transcoding"`
	CdnURL              string          `json:"cdn_url"`
	CurrentVersion      string          `json:"current_version"`
	LatestVersion       string          `json:"latest_version"`
	SongCount           int             `json:"song_count"`
	SongLength          float64         `json:"song_length"`
	QueueState          *QueueStateResp `json:"queue_state"`
}

// UserResp todo:: check if the type field is really needed
type UserResp struct {
	Avatar      string      `json:"avatar"`
	Email       string      `json:"email"`
	ID          int         `json:"id"`
	IsAdmin     bool        `json:"is_admin"`
	IsProspect  bool        `json:"is_prospect"`
	Name        string      `json:"name"`
	Preferences Preferences `json:"preferences"`
	Type        string      `json:"type"`
}

func WrapUser(user *model.User) *UserResp {
	return &UserResp{
		Avatar:      "https://www.gravatar.com/avatar/fbcba82c64cd973467a0e1b0600c1c1b?s=192&d=robohash",
		Email:       user.Email,
		ID:          user.ID,
		IsAdmin:     true,
		IsProspect:  false,
		Name:        user.Name,
		Preferences: Preferences{LastFmSessionKey: ""},
		Type:        "users",
	}
}

type Preferences struct {
	LastFmSessionKey interface{} `json:"lastfm_session_key"`
}

type QueueStateResp struct {
	*repository.QueueStateRepo `json:",inline"`
	Type                       string `json:"type"`
}

func WrapQueueState(state *repository.QueueStateRepo) *QueueStateResp {
	return &QueueStateResp{
		QueueStateRepo: state,
		Type:           "queue-states",
	}
}

type OverviewResp struct {
	MostPlayedAlbums    []*AlbumWrap    `json:"most_played_albums"`
	MostPlayedArtists   []*model.Artist `json:"most_played_artists"`
	MostPlayedSongs     []*model.Song   `json:"most_played_songs"`
	RecentlyAddedAlbums []*AlbumWrap    `json:"recently_added_albums"`
	RecentlyAddedSongs  []*SongWrap     `json:"recently_added_songs"`
	RecentlyPlayedSongs []*SongWrap     `json:"recently_played_songs"`
}

type AlbumWrap struct {
	*model.Album `json:",inline"`
	ArtistName   string `json:"artist_name"`
}

func WrapAlbums(albums []*model.Album) []*AlbumWrap {
	var result []*AlbumWrap
	for _, album := range albums {
		result = append(result, WrapAlbum(album))
	}
	return result
}

func WrapAlbum(album *model.Album) *AlbumWrap {
	return &AlbumWrap{
		Album:      album,
		ArtistName: album.Artist.Name,
	}
}

type SongWrap struct {
	AlbumArtistID   int     `json:"album_artist_id"`
	AlbumArtistName string  `json:"album_artist_name"`
	AlbumCover      string  `json:"album_cover"`
	AlbumID         int     `json:"album_id"`
	AlbumName       string  `json:"album_name"`
	ArtistID        int     `json:"artist_id"`
	ArtistName      string  `json:"artist_name"`
	CreatedAt       string  `json:"created_at"`
	Disc            int     `json:"disc"`
	Genre           string  `json:"genre"`
	ID              string  `json:"id"`
	Length          float64 `json:"length"`
	Liked           bool    `json:"liked"`
	Lyrics          string  `json:"lyrics"`
	PlayCount       int     `json:"play_count"`
	Title           string  `json:"title"`
	Track           int     `json:"track"`
	Type            string  `json:"type"`
	Year            int     `json:"year"`
}

func WrapSongs(songs []*model.Song) []*SongWrap {
	var result []*SongWrap
	for _, song := range songs {
		result = append(result, WrapSong(song))
	}
	return result
}

func WrapSong(song *model.Song) *SongWrap {
	liked, playCount := false, 0
	if song.Interaction != nil {
		liked = song.Interaction.Liked
		playCount = song.Interaction.PlayCount
	}

	return &SongWrap{
		AlbumArtistID:   song.Album.ArtistID,
		AlbumArtistName: song.Artist.Name,
		AlbumCover:      song.Album.Cover,
		AlbumID:         song.AlbumID,
		AlbumName:       song.Album.Name,
		ArtistID:        song.ArtistID,
		ArtistName:      song.Artist.Name,
		CreatedAt:       song.CreatedAt.String(),
		Disc:            song.Disc,
		Genre:           song.Genre,
		ID:              song.ID,
		Length:          song.Length,
		Liked:           liked,
		Lyrics:          song.Lyrics,
		PlayCount:       playCount,
		Title:           song.Title,
		Track:           song.Track,
		Year:            song.Year,
	}
}
