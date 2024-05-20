package repository

import (
	"context"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Artists struct {
	ctx context.Context
}

func NewArtists(ctx context.Context) *Artists {
	return &Artists{ctx: ctx}
}

func (a *Artists) FindOrCreate(artist *model.Artist) (*model.Artist, error) {
	existsArtist, err := dao.Artist.WithContext(a.ctx).Where(dao.Artist.Name.Eq(artist.Name)).First()

	if err != nil {
		err := dao.Artist.WithContext(a.ctx).Create(artist)
		return artist, err
	}
	return existsArtist, nil
}
