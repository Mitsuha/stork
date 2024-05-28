package playlists

type CreateReq struct {
	Name  string   `json:"name"`
	Songs []string `json:"songs"`
}

type UpdateReq struct {
	Name string `json:"name"`
}

func (u *UpdateReq) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": u.Name,
	}
}

type PlaylistSongsReq struct {
	ID int `uri:"id"`
}
