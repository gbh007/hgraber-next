package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"

	"hgnext/internal/config"
)

func parseConfig() (config.Config, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	c := config.ConfigDefault()

	f, err := os.Open(*configPath)
	if err != nil {
		return config.Config{}, fmt.Errorf("open config file: %w", err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return config.Config{}, fmt.Errorf("decode yaml: %w", err)
	}

	err = envconfig.Process("APP", &c)
	if err != nil {
		return config.Config{}, fmt.Errorf("decode env: %w", err)
	}

	return c, nil
}
