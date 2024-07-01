package users

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/repository/model/dao"
)

type User struct {
}

func New() *User {
	return &User{}
}

func (u *User) Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	user, err := authentication.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	resp := LoginResp{
		Token: authentication.NewAccessToken(user, []string{"*"}),
		Audio: authentication.NewAccessToken(user, []string{"audio"}),
	}

	if err := dao.PersonalAccessToken.WithContext(ctx).Create(resp.Token.Model, resp.Audio.Model); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, resp)
}

func (u *User) Index(ctx *gin.Context) {
	users, err := dao.User.WithContext(ctx).Find()
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, WrapUserSummary(users))
}
