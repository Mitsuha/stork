package ipfs

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	iface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo"
	"github.com/ipfs/kubo/repo/fsrepo"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func Run(ctx context.Context, cfg *Config) (iface.CoreAPI, *core.IpfsNode, error) {
	if cfg.RepoPath != "" {
		rep, err := createRepo(cfg.RepoPath)
		if err != nil {
			return nil, nil, nil
		}
		cfg.Repo = rep
	}
	if cfg.Repo == nil {
		return nil, nil, errors.New("no repo provided")
	}

	return createNode(ctx, &cfg.BuildCfg)
}

var loadPluginsOnce sync.Once

func createRepo(p string) (repo.Repo, error) {
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return nil, onceErr
	}

	if err := os.Mkdir(p, 0755); err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}

	cfg, err := config.Init(io.Discard, 2048)
	if err != nil {
		return nil, err
	}

	if err := fsrepo.Init(p, cfg); err != nil {
		return nil, fmt.Errorf("failed to initialize the repo: %s", err)
	}

	return fsrepo.Open(p)
}

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

func createNode(ctx context.Context, cfg *core.BuildCfg) (iface.CoreAPI, *core.IpfsNode, error) {
	node, err := core.NewNode(ctx, cfg)

	if err != nil {
		return nil, nil, err
	}

	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, nil, err
	}

	return api, node, err
}
