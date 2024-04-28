package main

import (
	"flag"
	"fmt"
	"github.com/mitsuha/stork/config"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/repository/model"
	"gorm.io/gen"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to the configuration file.")
	flag.Parse()

	if err := config.Load(configPath); err != nil {
		fmt.Println("Failed to load the configuration: ", err)
		return
	}

	if err := container.Boot(); err != nil {
		fmt.Println("Failed to boot the container: ", err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./repository/model/dao",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(container.Singled.DB)

	model.ApplyQueries(g)

	g.Execute()
}
