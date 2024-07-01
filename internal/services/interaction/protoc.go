package interaction

type SongReq struct {
	Song string `json:"song" validate:"required"`
}
