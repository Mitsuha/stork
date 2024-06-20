package songs

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ipfs/boxo/files"
	"github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/pkg/audio"
	"github.com/mitsuha/stork/pkg/authentication"
	"github.com/mitsuha/stork/pkg/ipfs"
	"github.com/mitsuha/stork/pkg/paginate"
	"github.com/mitsuha/stork/repository"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"io"
	"net/url"
	"path/filepath"
	"strings"
	"sync"
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
	var req UploadReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	file, err := req.File.Open()
	if err != nil {
		ctx.JSON(500, errors.Join(v1.ServerError, err))
		return
	}

	metadata, err := audio.New(audio.OSPath).LoadFromReader(file).Metadata()
	if err != nil {
		ctx.JSON(500, errors.Join(v1.ServerError, err))
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(500, errors.Join(v1.ServerError, err))
		return
	}

	path, err := ipfs.Unixfs().Add(ctx, files.NewReaderFile(file))
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
			Path:   path.RootCid().String(),
			Album:  album,
			Artist: artist,
		}
	} else {
		title, _ := url.QueryUnescape(req.File.Filename)

		title = strings.TrimSuffix(title, filepath.Ext(title))

		song = &model.Song{
			Title:  title,
			Length: metadata.Duration(),
			Path:   path.RootCid().String(),
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

var fsPool = sync.Pool{
	New: func() any {
		return ipfs.NewFilesystem(nil, nil)
	},
}

func SetupFS(ctx context.Context) func(filesystem *ipfs.Filesystem) {
	return func(f *ipfs.Filesystem) {
		f.IPFS = ipfs.Instance()
		f.Ctx = ctx
	}
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

	fs := fsPool.Get().(*ipfs.Filesystem)
	defer func() {
		fsPool.Put(fs)
	}()

	fs.WithOption(SetupFS(ctx))

	ctx.FileFromFS(song.Path, fs)
}
