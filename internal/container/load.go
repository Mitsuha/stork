package container

import (
	"context"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/mitsuha/stork/config"
	"github.com/mitsuha/stork/pkg/ipfs"
	"github.com/mitsuha/stork/repository/model/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Boot(ctx context.Context) error {
	db, err := gorm.Open(mysql.Open(config.Mysql.ToDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	api, node, err := ipfs.Run(ctx, &ipfs.Config{
		BuildCfg: core.BuildCfg{
			Online:  true,
			Routing: libp2p.DHTOption,
		},
		RepoPath: ".ipfs",
	})
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		_ = node.Close()
	}()

	Singled = &singleton{
		DB:       db,
		Ipfs:     api,
		IpfsNode: node,
	}

	dao.SetDefault(db)

	return nil
}
