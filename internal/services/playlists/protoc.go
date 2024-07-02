package playlists

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlaylistReq struct {
	ID int `uri:"id"`
}

type CreateReq struct {
	Name  string   `json:"name"`
	Songs []string `json:"songs"`
}

// UpdateReq When it is used for updating the name of the playlist.s
type UpdateReq struct {
	ID   int    `uri:"id"`
	Name string `json:"name" form:"name"`
}

func (u *UpdateReq) Bind(ctx *gin.Context) (err error) {
	if err = ctx.BindJSON(u); err != nil {
		return
	}

	u.ID, err = strconv.Atoi(ctx.Param("id"))
	return
}

func (u *UpdateReq) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": u.Name,
	}
}

// AssociateActionReq When it is used for adding or deleting songs.
type AssociateActionReq struct {
	ID    int      `uri:"id"`
	Songs []string `json:"songs" form:"song"`
}

func (a *AssociateActionReq) Bind(ctx *gin.Context) (err error) {
	if err = ctx.BindJSON(a); err != nil {
		return
	}

	a.ID, err = strconv.Atoi(ctx.Param("id"))
	return
}
