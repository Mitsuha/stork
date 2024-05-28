package artists

type ShowReq struct {
	ID int `json:"id" uri:"id" binding:"required"`
}
