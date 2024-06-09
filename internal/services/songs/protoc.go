package songs

import (
	"github.com/gin-gonic/gin"
	"github.com/mitsuha/stork/pkg/paginate"
	"github.com/mitsuha/stork/pkg/slices"
	"mime/multipart"
)

type IndexReq struct {
	paginate.Request
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

var optionalColumn = []string{"title", "track", "time", "created_at"}
var optionalOrder = []string{"asc", "desc"}

func (i *IndexReq) BindRequest(ctx *gin.Context) error {
	if err := ctx.BindQuery(i); err != nil {
		return err
	}
	if i.Page == 0 {
		i.Page = 1
	}
	if i.PageSize == 0 {
		i.PageSize = paginate.PageSize
	}

	if !slices.Container(optionalColumn, i.Sort) {
		i.Sort = optionalColumn[0]
	}

	if !slices.Container(optionalOrder, i.Order) {
		i.Order = optionalOrder[0]
	}
	return nil
}

type PlayReq struct {
	ID string `uri:"id" `
}

type UploadReq struct {
	File *multipart.FileHeader `form:"file" binding:"required,audioOnly"`
}
