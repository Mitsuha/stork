package queue

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/repository"
)

type Queue struct {
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) UpdateState(ctx *gin.Context) {
	user := authentication.User(ctx)
	var req UpdateQueueStateReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	if err := repository.NewQueueStates(ctx).UpdateQueueState(user.ID, req.Songs); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
}

func (q *Queue) PlaybackStatus(ctx *gin.Context) {
	user := authentication.User(ctx)

	var req UpdatePlaybackStatusReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	if err := repository.NewQueueStates(ctx).UpdatePlayState(user.ID, req.Song, req.Position); err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}
}
