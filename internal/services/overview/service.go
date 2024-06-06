package overview

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/repository"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Overview struct {
}

func New() *Overview {
	return &Overview{}
}

func (o *Overview) Data(ctx *gin.Context) {
	user := authentication.User(ctx)

	state, err := repository.NewQueueStates(ctx).UsersQueueState(user.ID)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
	statist, err := dao.Song.WithContext(ctx).CountAndLength()
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	playlists, err := dao.Playlist.WithContext(ctx).FindByUserID(user.ID)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, &DataResponse{
		Settings:            repository.Settings(),
		Playlists:           WrapPlaylist(playlists),
		PlaylistFolders:     []string{},
		CurrentUser:         WrapUser(user),
		UseLastFm:           false,
		UseSpotify:          false,
		UseYouTube:          false,
		UseITunes:           false,
		AllowDownload:       true,
		SupportsTranscoding: false,
		CdnURL:              "http://localhost:8080/",
		CurrentVersion:      "v6.12.1",
		LatestVersion:       "v6.12.1",
		SongCount:           statist.Count,
		SongLength:          statist.Length,
		QueueState:          WrapQueueState(state),
	})
}

func (o *Overview) Overview(ctx *gin.Context) {
	user := authentication.User(ctx)

	albumMostPlayed, err := dao.Album.WithContext(ctx).Preload(dao.Album.Artist).MostPlayed(user.ID, 5)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	artistMostPlayed, err := dao.Artist.WithContext(ctx).MostPlayed(user.ID, 5)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	songsMostPlayed, err := dao.Song.WithContext(ctx).MostPlayed(user.ID, 5)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	songsRecentlyPlayed, err := dao.Song.WithContext(ctx).Preload(dao.Song.Album, dao.Song.Artist, dao.Song.Interaction).RecentlyPlayed(user.ID, 7)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	albumRecentlyAdded, err := dao.Album.WithContext(ctx).Preload(dao.Album.Artist).RecentlyAdded(5)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	songRecentlyAdded, err := dao.Song.WithContext(ctx).Preload(dao.Song.Album, dao.Song.Artist, dao.Song.Interaction).RecentlyAdded(user.ID, 7)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, &OverviewResp{
		MostPlayedAlbums:    WrapAlbums(albumMostPlayed),
		MostPlayedArtists:   artistMostPlayed,
		MostPlayedSongs:     songsMostPlayed,
		RecentlyPlayedSongs: WrapSongs(songsRecentlyPlayed),
		RecentlyAddedAlbums: WrapAlbums(albumRecentlyAdded),
		RecentlyAddedSongs:  WrapSongs(songRecentlyAdded),
	})
}
