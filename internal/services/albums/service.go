package albums

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/pkg/paginate"
	"github.com/mitsuha/stork/repository/model"
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

func (a *Albums) Index(ctx *gin.Context) {
	var req IndexReq

	if err := req.BindRequest(ctx); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	query := container.Singled.DB.Order("name ASC")

	page, err := paginate.Simple[model.Album](query, req.Request)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, page)
}

func (a *Albums) Songs(ctx *gin.Context) {
	var req ShowReq

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	fmt.Println(req.ID)

	songs, err := dao.Song.WithContext(ctx).
		Preload(dao.Song.Album, dao.Song.Artist, dao.Song.Interaction).
		Where(dao.Song.AlbumID.Eq(req.ID)).
		Order(dao.Song.Disc, dao.Song.Track, dao.Song.Title).
		Find()

	if err != nil {
		ctx.JSON(404, v1.NotFound)
		return
	}

	ctx.JSON(200, overview.WrapSongs(songs))

}
