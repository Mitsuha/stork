package gateway

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "github.com/mitsuha/stork/api/v1"
	"github.com/mitsuha/stork/pkg/ipfs"
	"sync"
	"time"
)

type IPFSGateway struct {
	fsPool sync.Pool
}

func New() *IPFSGateway {
	return &IPFSGateway{
		fsPool: sync.Pool{
			New: func() any {
				return ipfs.NewFilesystem(nil, nil)
			},
		},
	}
}

func SetupFS(ctx context.Context) func(filesystem *ipfs.Filesystem) {
	return func(f *ipfs.Filesystem) {
		f.IPFS = ipfs.Instance()
		f.Ctx = ctx
	}
}

func (i *IPFSGateway) File(ctx *gin.Context) {
	var req FileRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(400, v1.BadRequest)
		return
	}

	fs := i.fsPool.Get().(*ipfs.Filesystem)
	tc, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer func() {
		i.fsPool.Put(fs)
		cancel()
	}()

	fs.WithOption(SetupFS(tc))

	ctx.FileFromFS(req.Filename, fs)
}
