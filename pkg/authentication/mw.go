package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/repository/model"
	"strings"
)

const (
	AuthHeader    = "Authorization"
	AuthPrefix    = "Bearer "
	AuthUserInCtx = "auth:user"
)

func Auth(ctx *gin.Context) {
	token := strings.Replace(ctx.GetHeader(AuthHeader), AuthPrefix, "", 1)

	user, err := getUser(ctx, token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, v1.Unauthorized)
		return
	}

	ctx.Set(AuthUserInCtx, user)

	ctx.Next()
}

func LazyAuth(ctx *gin.Context) {
	token := strings.Replace(ctx.GetHeader(AuthHeader), AuthPrefix, "", 1)

	user, _ := getUser(ctx, token)

	ctx.Set(AuthUserInCtx, user)

	ctx.Next()
}

func User(ctx *gin.Context) *model.User {
	return ctx.MustGet(AuthUserInCtx).(*model.User)
}
