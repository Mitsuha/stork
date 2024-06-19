package container

import (
	"github.com/ipfs/kubo/core"
	iface "github.com/ipfs/kubo/core/coreiface"
	"gorm.io/gorm"
)

var Singled *singleton

type singleton struct {
	DB       *gorm.DB
	Ipfs     iface.CoreAPI
	IpfsNode *core.IpfsNode
}
