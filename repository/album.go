package repository

import (
	"context"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type Albums struct {
	ctx context.Context
}

func NewAlbums(ctx context.Context) *Albums {
	return &Albums{ctx: ctx}
}

func (a *Albums) FindOrCreate(album *model.Album) (*model.Album, error) {
	existsAlbum, err := dao.Album.WithContext(a.ctx).Where(dao.Album.Name.Eq(album.Name)).Where(dao.Album.ArtistID.Eq(album.ArtistID)).First()

	if err != nil {
		err := dao.Album.WithContext(a.ctx).Create(album)
		return album, err
	}
	if existsAlbum.Cover == "" && album.Cover != "" {
		existsAlbum.Cover = album.Cover
		_, _ = dao.Album.WithContext(a.ctx).Where(dao.Album.ID.Eq(existsAlbum.ID)).Update(dao.Album.Cover, album.Cover)
	}

	return existsAlbum, nil
}
