package albums

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Albums struct {
}

func New() *Albums {
	return &Albums{}
}

func (a *Albums) Show(ctx *gin.Context) {
	var req ShowReq

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	album, err := dao.Album.WithContext(ctx).Preload(dao.Album.Artist).FindByID(req.ID)
	if err != nil {
		ctx.JSON(404, v1.NotFound)
		return
	}

	ctx.JSON(200, WrapAlbum(album))
}

func (a *Albums) Songs(ctx *gin.Context) {
	var req ShowReq

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	fmt.Println(req.ID)

	songs, err := dao.Songs.WithContext(ctx).
		Preload(dao.Songs.Album, dao.Songs.Artist, dao.Songs.Interaction).
		Where(dao.Songs.AlbumID.Eq(req.ID)).
		Order(dao.Songs.Disc, dao.Songs.Track, dao.Songs.Title).
		Find()

	if err != nil {
		ctx.JSON(404, v1.NotFound)
		return
	}

	ctx.JSON(200, overview.WrapSongs(songs))

}
