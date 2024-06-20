package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/mitsuha/stork/config"
	"github.com/mitsuha/stork/internal"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/pkg/ipfs"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to the configuration file.")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	go listenCancel(cancel)

	if err := config.Load(configPath); err != nil {
		fmt.Println("Failed to load the configuration: ", err)
		return
	}

	if err := container.Boot(); err != nil {
		fmt.Println("Failed to boot the container: ", err)
		return
	}

	if err := ipfs.Start(ctx, defaultIPFSConfig()); err != nil {
		panic(err)
	}

	if err := internal.Run(); err != nil {
		panic(err)
	}
}

func listenCancel(cancel context.CancelFunc) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

	<-sigint
	cancel()
	_ = ipfs.Close()
	os.Exit(0)
}

func defaultIPFSConfig() *ipfs.Config {
	return &ipfs.Config{
		BuildCfg: core.BuildCfg{
			Online:  true,
			Routing: libp2p.DHTOption,
		},
		RepoPath: ".ipfs",
	}
}
