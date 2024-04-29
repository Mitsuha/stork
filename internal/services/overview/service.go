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

	state, err := repository.UsersQueueState(ctx, user.ID)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
	statist, err := dao.Songs.WithContext(ctx).CountAndLength()
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, &DataResponse{
		Settings:            repository.Settings(),
		Playlists:           "",
		PlaylistFolders:     "",
		CurrentUser:         WrapUser(user),
		UseLastFm:           false,
		UseSpotify:          false,
		UseYouTube:          false,
		UseITunes:           false,
		AllowDownload:       true,
		SupportsTranscoding: false,
		CdnURL:              "http://koel.test",
		CurrentVersion:      "v6.12.1",
		LatestVersion:       "v6.12.1",
		SongCount:           statist.Count,
		SongLength:          statist.Length,
		QueueState:          WrapQueueState(state),
	})
}
