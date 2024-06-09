package songs

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/pkg/audio"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/pkg/paginate"
	"github.com/mitsuha/stork/repository"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"net/url"
	"path/filepath"
	"strings"
)

type Songs struct {
}

const (
	RecentlyPlayedCount = 128
)

var (
	unknownArtist = &model.Artist{Name: "Unknown Artist"}
	unknownAlbum  = &model.Album{Name: "Unknown Album"}
)

func New() *Songs {
	return &Songs{}
}

func (s *Songs) Index(ctx *gin.Context) {
	var req IndexReq
	if err := req.BindRequest(ctx); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	query := container.Singled.DB.Order(fmt.Sprintf("%s %s", req.Sort, req.Order)).Preload("Album").Preload("Artist").Preload("Interaction")

	page, err := paginate.Simple[*model.Song](query, req.Request)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, paginate.Page[*overview.SongWrap]{
		Data: overview.WrapSongs(page.Data),
		Meta: page.Meta,
	})
}

func (s *Songs) Upload(ctx *gin.Context) {
	settings := repository.Settings()

	var req UploadReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	path := fmt.Sprintf("%s/uploads/%s", settings.MediaPath(), req.File.Filename)

	if err := ctx.SaveUploadedFile(req.File, path); err != nil {
		ctx.JSON(500, errors.Join(v1.ServerError, err))
		return
	}

	file, _ := req.File.Open()

	metadata, err := audio.New(audio.OSPath).LoadFromReader(file).Metadata()
	if err != nil {
		ctx.JSON(500, errors.Join(v1.ServerError, err))
		return
	}

	var artist, album = unknownArtist, unknownAlbum
	var song *model.Song

	if metadata.HasTag() {
		artist = &model.Artist{
			Name: metadata.Artist(),
		}
		album = &model.Album{
			Name:  metadata.Album(),
			Cover: "",
		}

		song = &model.Song{
			ID:     uuid.NewString(),
			Title:  metadata.Title(),
			Year:   metadata.Year(),
			Genre:  metadata.Genre(),
			Track:  metadata.Track(),
			Disc:   metadata.Track(),
			Length: metadata.Duration(),
			Path:   path,
			Album:  album,
			Artist: artist,
		}
	} else {
		title, _ := url.QueryUnescape(req.File.Filename)

		title = strings.TrimSuffix(title, filepath.Ext(title))

		song = &model.Song{
			Title:  title,
			Length: metadata.Duration(),
			Path:   path,
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

func (s *Songs) RecentlyPlayed(ctx *gin.Context) {
	user := authentication.User(ctx)

	songs, err := dao.Song.WithContext(ctx).Preload(dao.Song.Album, dao.Song.Artist, dao.Song.Interaction).RecentlyPlayed(user.ID, RecentlyPlayedCount)
	if err != nil {
		ctx.JSON(500, v1.ServerError)
		return
	}

	ctx.JSON(200, overview.WrapSongs(songs))
}

func (s *Songs) Play(ctx *gin.Context) {
	var req PlayReq
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}
	song, err := dao.Song.WithContext(ctx).FindByID(req.ID)
	if err != nil {
		ctx.JSON(404, v1.NotFound)
		return
	}

	ctx.File(song.Path)
}
