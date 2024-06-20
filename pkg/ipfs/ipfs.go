package ipfs

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/boxo/path"
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

var coreApi iface.CoreAPI

var node *core.IpfsNode

func Start(ctx context.Context, cfg *Config) (err error) {
	if cfg.RepoPath != "" {
		rep, err := createRepo(cfg.RepoPath)
		if err != nil {
			return err
		}
		cfg.Repo = rep
	}
	if cfg.Repo == nil {
		return errors.New("no repo provided")
	}

	coreApi, node, err = createNode(ctx, &cfg.BuildCfg)

	return err
}

func Close() error {
	defer func() {
		coreApi, node = nil, nil
	}()

	return node.Close()
}

func Instance() iface.CoreAPI {
	return coreApi
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

// Unixfs returns an implementation of Unixfs API
func Unixfs() iface.UnixfsAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Unixfs()
}

// Block returns an implementation of Block API
func Block() iface.BlockAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Block()
}

// Dag returns an implementation of Dag API
func Dag() iface.APIDagService {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Dag()
}

// Name returns an implementation of Name API
func Name() iface.NameAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Name()
}

// Pin returns an implementation of Pin API
func Pin() iface.PinAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Pin()
}

// Key returns an implementation of Key API
func Key() iface.KeyAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Key()
}

// ResolvePath returns an implementation of ResolvePath API
func ResolvePath(ctx context.Context, p path.Path) (path.ImmutablePath, []string, error) {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.ResolvePath(ctx, p)
}

// Swarm returns an implementation of Swarm API
func Swarm() iface.SwarmAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Swarm()
}

// Object returns an implementation of Object API
func Object() iface.ObjectAPI {
	if coreApi == nil || node == nil {
		panic("ipfs not started")
	}
	return coreApi.Object()
}
