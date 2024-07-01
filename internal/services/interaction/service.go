package interaction

import (
	"github.com/gin-gonic/gin"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/pkg/ipfs"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"time"
)

type Interaction struct {
}

func New() *Interaction {
	return &Interaction{}
}

func (i *Interaction) Play(ctx *gin.Context) {
	user := authentication.User(ctx)

	var req SongReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	interaction, err := dao.Interaction.WithContext(ctx).FindByUserSong(user.ID, req.Song)
	if err != nil {
		interaction = &model.Interaction{
			UserID:       user.ID,
			SongID:       req.Song,
			PlayCount:    1,
			LastPlayedAt: time.Now(),
		}

		if err := dao.Interaction.WithContext(ctx).Create(interaction); err != nil {
			ctx.JSON(500, v1.ServerError)
			return
		}
		return
	}

	if err := dao.Interaction.WithContext(ctx).IncPlayCount(user.ID, req.Song); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
}

func (i *Interaction) ToggleLike(ctx *gin.Context) {
	user := authentication.User(ctx)

	var req SongReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	interaction, err := dao.Interaction.WithContext(ctx).FindByUserSong(user.ID, req.Song)
	if err != nil {
		interaction = &model.Interaction{
			UserID:       user.ID,
			SongID:       req.Song,
			LastPlayedAt: time.Now(),
		}

		if err := dao.Interaction.WithContext(ctx).Create(interaction); err != nil {
			ctx.JSON(500, v1.ServerError)
			return
		}
		return
	}
	interaction.Liked = !interaction.Liked

	if err := dao.Interaction.WithContext(ctx).ToggleLike(user.ID, req.Song, interaction.Liked); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	if interaction.Liked {
		return
	}

	has, _ := dao.Interaction.WithContext(ctx).HasOtherLiked(user.ID, req.Song)
	if !has {
		song, err := dao.Song.WithContext(ctx).FindByID(req.Song)
		if err != nil || (song.From == model.SFromLocal) {
			return
		}

		if c, err := cid.Parse(song.Path); err == nil {
			_ = ipfs.Pin().Rm(ctx, path.FromCid(c))
		}
	}
}
