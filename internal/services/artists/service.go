package artists

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/pkg/paginate"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Artists struct {
}

func New() *Artists {
	return &Artists{}
}

func (a *Artists) Index(ctx *gin.Context) {
	var req IndexReq
	if err := req.BindRequest(ctx); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	query := container.Singled.DB.Order("name ASC")

	page, err := paginate.Simple[model.Artist](query, req.Request)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
	ctx.JSON(200, page)
}

func (a *Artists) Show(ctx *gin.Context) {
	var req ShowReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	artist, err := dao.Artist.WithContext(ctx).FindByID(req.ID)
	if err != nil {
		ctx.JSON(500, v1.NotFound)
		return
	}

	ctx.JSON(200, artist)
}
