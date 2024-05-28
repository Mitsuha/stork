package songs

import (
	"errors"
	"fmt"
	"github.com/dhowden/tag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/repository"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"net/url"
	"path/filepath"
	"strings"
)

type Songs struct {
}

var (
	unknownArtist = &model.Artist{Name: "Unknown Artist"}
	unknownAlbum  = &model.Album{Name: "Unknown Album"}
)

func New() *Songs {
	return &Songs{}
}

func (s *Songs) Upload(ctx *gin.Context) {
	var req UploadReq

	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(400, v1.BadRequest)
		return
	}
	err := ctx.SaveUploadedFile(req.File, fmt.Sprintf("uploads/%s", req.File.Filename))

	file, _ := req.File.Open()

	metadata, err := tag.ReadFrom(file)

	fmt.Println(metadata.Raw())

	var artist, album = unknownArtist, unknownAlbum
	var song *model.Song

	if err == nil {
		artist = &model.Artist{
			Name: metadata.Artist(),
		}
		album = &model.Album{
			Name:  metadata.Album(),
			Cover: "",
		}

		_, track := metadata.Track()
		_, disc := metadata.Disc()

		song = &model.Song{
			ID:    uuid.NewString(),
			Title: metadata.Title(),
			Year:  metadata.Year(),
			Genre: metadata.Genre(),
			Track: track,
			Disc:  disc,
			//Length: metadata.,
			Path:   fmt.Sprintf("uploads/%s", req.File.Filename),
			Album:  album,
			Artist: artist,
		}
	} else {
		title, _ := url.QueryUnescape(req.File.Filename)

		title = strings.TrimSuffix(title, filepath.Ext(title))

		song = &model.Song{
			Title: title,
			//Length: metadata.,
			Path:   fmt.Sprintf("uploads/%s", req.File.Filename),
			Album:  album,
			Artist: artist,
		}
	}

	if err := repository.NewSongs(ctx).Create(song); err != nil {
		ctx.JSON(500, errors.Join(v1.ServerError, err))
		return
	}

	ctx.JSON(200, gin.H{
		"album": overview.WrapAlbum(album),
		"song":  overview.WrapSong(song),
	})
}

func (s *Songs) Favorite(ctx *gin.Context) {
	user := authentication.User(ctx)

	songs, err := dao.Song.WithContext(ctx).Preload(dao.Song.Album, dao.Song.Artist, dao.Song.Interaction).Favorite(user.ID)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, overview.WrapSongs(songs))
}
