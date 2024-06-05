package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mitsuha/stork/repository/model/dao"
)

type User struct {
}

func New() *User {
	return &User{}
}

func (u *User) Index(ctx *gin.Context) {
	users, err := dao.User.WithContext(ctx).Find()
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, WrapUserSummary(users))
}
