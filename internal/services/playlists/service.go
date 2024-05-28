package playlists

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Playlist struct {
}

func New() *Playlist {
	return &Playlist{}
}

func (p *Playlist) Create(ctx *gin.Context) {
	user := authentication.User(ctx)

	var req CreateReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	exits, _ := dao.Song.WithContext(ctx).Where(dao.Song.ID.In(req.Songs...)).Count()
	if int(exits) != len(req.Songs) {
		ctx.JSON(400, v1.BadRequest)
	}

	playlist := model.Playlist{
		UserID: user.ID,
		Name:   req.Name,
	}

	if err := dao.Playlist.WithContext(ctx).Create(&playlist); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
	ps := make([]*model.PlaylistSong, len(req.Songs))
	for i, song := range req.Songs {
		ps[i] = &model.PlaylistSong{
			PlaylistID: playlist.ID,
			SongID:     song,
		}
	}

	if err := dao.PlaylistSong.WithContext(ctx).Create(ps...); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, playlist)
}

func (p *Playlist) Songs(ctx *gin.Context) {
	var req PlaylistSongsReq
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	songs, err := dao.Song.WithContext(ctx).Preload(dao.Song.Album, dao.Song.Artist, dao.Song.Interaction).FindByPlaylist(req.ID)
	if err != nil {
		return
	}

	ctx.JSON(200, overview.WrapSongs(songs))
}

func (p *Playlist) Update(ctx *gin.Context) {
	user := authentication.User(ctx)

	var req UpdateReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}
	var idReq PlaylistSongsReq
	if err := ctx.BindUri(&idReq); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	playlist, err := dao.Playlist.WithContext(ctx).FindByID(idReq.ID)
	if err != nil {
		ctx.JSON(500, v1.NotFound)
		return
	}
	if playlist.UserID != user.ID {
		ctx.JSON(403, v1.Forbidden)
		return
	}

	_, err = dao.Playlist.WithContext(ctx).Where(dao.Playlist.ID.Eq(idReq.ID)).Updates(req.ToMap())
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	playlist.Name = req.Name

	ctx.JSON(200, playlist)
}
