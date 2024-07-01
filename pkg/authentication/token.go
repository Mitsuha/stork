package authentication

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/mitsuha/stork/pkg/strs"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"strings"
)

type AccessToken struct {
	Model      *model.PersonalAccessToken `json:"-"`
	PlainToken string                     `json:"token"`
}

func (a *AccessToken) Token() string {
	return fmt.Sprintf("%d|%s", a.Model.ID, a.PlainToken)
}

func NewAccessToken(user *model.User, abilities []string) *AccessToken {
	plainTextToken := strs.Random(40)

	s := sha256.New()
	s.Write([]byte(plainTextToken))

	return &AccessToken{
		Model: &model.PersonalAccessToken{
			ID:          0,
			TokenableID: user.ID,
			Name:        user.Name,
			Token:       hex.EncodeToString(s.Sum(nil)),
			Abilities:   abilities,
		},
		PlainToken: plainTextToken,
	}
}

func accessTokenFromEncodingToken(ctx context.Context, et string) (*AccessToken, error) {
	t := strings.Split(et, "|")
	if len(t) != 2 {
		return nil, errors.New("invalid token")
	}

	h := sha256.New()

	h.Write([]byte(t[1]))

	pt, err := dao.PersonalAccessToken.WithContext(ctx).WhereIDAndToken(t[0], hex.EncodeToString(h.Sum(nil)))

	if err != nil {
		return nil, err
	}

	return &AccessToken{Model: pt, PlainToken: t[1]}, nil
}
