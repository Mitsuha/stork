package paginate

import "github.com/gin-gonic/gin"

const (
	PageSize = 15
)

type Request struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize"`
}

func (i *Request) BindRequest(ctx *gin.Context) error {
	if err := ctx.BindQuery(i); err != nil {
		return err
	}
	if i.Page == 0 {
		i.Page = 1
	}
	if i.PageSize == 0 {
		i.PageSize = PageSize
	}
	return nil
}
