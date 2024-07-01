package users

import (
	"github.com/goccy/go-json"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/repository/model"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token *authentication.AccessToken `json:"api"`
	Audio *authentication.AccessToken `json:"audio"`
}

func (l LoginResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"token":       l.Token.Token(),
		"audio-token": l.Audio.Token(),
	})
}

type UserSummary struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
	IsAdmin    bool   `json:"is_admin"`
	IsProspect bool   `json:"is_prospect"`
}

func WrapUserSummary(users []*model.User) []UserSummary {
	data := make([]UserSummary, len(users))

	for i, user := range users {
		data[i] = UserSummary{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Avatar:     "",
			IsAdmin:    user.IsAdmin,
			IsProspect: false,
		}
	}

	return data
}
