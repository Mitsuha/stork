package authentication

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

type UserWrap struct {
	model.User `json:",inline"`

	Token TokenWrap `json:"token"`
}

type TokenWrap struct {
	Model *model.PersonalAccessToken `json:"-"`
	Token string                     `json:"token"`
}

func NewWrap(model *model.PersonalAccessToken) *TokenWrap {
	h := sha3.New256()
	h.Write([]byte(model.Token))

	return &TokenWrap{
		Model: model,
		Token: strconv.Itoa(model.ID) + "|" + string(h.Sum(nil)),
	}
}

func getUser(ctx context.Context, token string) (*model.User, error) {
	wrap, err := getToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return dao.User.WithContext(ctx).FindByID(wrap.Model.TokenableID)
}

func getToken(ctx context.Context, token string) (*TokenWrap, error) {
	t := strings.Split(token, "|")
	if len(t) != 2 {
		return nil, errors.New("invalid token")
	}

	h := sha256.New()

	h.Write([]byte(t[1]))

	tm, err := dao.PersonalAccessToken.WithContext(ctx).WhereIDAndToken(t[0], hex.EncodeToString(h.Sum(nil)))

	if err != nil {
		return nil, err
	}

	return &TokenWrap{Model: tm}, nil
}
