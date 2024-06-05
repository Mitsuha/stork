package users

import "github.com/mitsuha/stork/repository/model"

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
