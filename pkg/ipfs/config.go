package ipfs

import "github.com/ipfs/kubo/core"

type Config struct {
	core.BuildCfg
	RepoPath string
}
