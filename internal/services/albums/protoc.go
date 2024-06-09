package albums

import (
	"github.com/mitsuha/stork/pkg/paginate"
	"github.com/mitsuha/stork/repository/model"
)

type IndexReq struct {
	paginate.Request
}

type ShowReq struct {
	ID int `json:"id" uri:"id" binding:"required"`
}

type AlbumWrap struct {
	*model.Album `json:",inline"`
	ArtistName   string `json:"artist_name"`
}

func WrapAlbum(album *model.Album) *AlbumWrap {
	return &AlbumWrap{
		Album:      album,
		ArtistName: album.Artist.Name,
	}
}
