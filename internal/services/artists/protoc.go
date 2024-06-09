package artists

import "github.com/mitsuha/stork/pkg/paginate"

type IndexReq struct {
	paginate.Request
}

type ShowReq struct {
	ID int `json:"id" uri:"id" binding:"required"`
}
