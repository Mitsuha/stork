package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mitsuha/stork/config"
	"github.com/mitsuha/stork/internal"
	"github.com/mitsuha/stork/internal/container"
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

	if err := container.Boot(ctx); err != nil {
		fmt.Println("Failed to boot the container: ", err)
		return
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
	os.Exit(0)
}
