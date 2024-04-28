package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type config struct {
	Mysql *mysql `json:"mysql" yaml:"mysql"`
}

func Load(path string) error {
	content, err := getFileContent(path)
	if err != nil {
		return err
	}

	var c config
	if err := yaml.Unmarshal([]byte(os.ExpandEnv(content)), &c); err != nil {
		return err
	}

	overwrite(&c)

	return nil
}

func getFileContent(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(all), nil
}

func overwrite(c *config) {
	Mysql = c.Mysql
}
