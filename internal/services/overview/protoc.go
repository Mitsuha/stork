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
