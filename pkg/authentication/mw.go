package authentication

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"strings"
)

const (
	AuthHeader    = "Authorization"
	AuthPrefix    = "Bearer "
	AuthUserInCtx = "auth:user"
)

func Auth(ctx *gin.Context) {
	token := strings.Replace(ctx.GetHeader(AuthHeader), AuthPrefix, "", 1)

	user, err := userFromToken(ctx, token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, v1.Unauthorized)
		return
	}

	ctx.Set(AuthUserInCtx, user)

	ctx.Next()
}

func LazyAuth(ctx *gin.Context) {
	token := strings.Replace(ctx.GetHeader(AuthHeader), AuthPrefix, "", 1)

	user, _ := userFromToken(ctx, token)

	ctx.Set(AuthUserInCtx, user)

	ctx.Next()
}

func userFromToken(ctx context.Context, token string) (*model.User, error) {
	wrap, err := accessTokenFromEncodingToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return dao.User.WithContext(ctx).FindByID(wrap.Model.TokenableID)
}

func User(ctx *gin.Context) *model.User {
	return ctx.MustGet(AuthUserInCtx).(*model.User)
}
