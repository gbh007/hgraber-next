package config

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func ImportConfig[T any](filename string, useEnv bool, newConfig func() T) (T, error) {
	var cfg T

	if newConfig != nil {
		cfg = newConfig()
	}

	ext := path.Ext(filename)

	switch ext {
	case ".yml", ".yaml":
		f, err := os.Open(filename)
		if err != nil {
			return cfg, fmt.Errorf("open config file: %w", err)
		}

		defer f.Close()

		err = yaml.NewDecoder(f).Decode(&cfg)
		if err != nil {
			return cfg, fmt.Errorf("decode yaml: %w", err)
		}
	case ".toml":
		f, err := os.Open(filename)
		if err != nil {
			return cfg, fmt.Errorf("open config file: %w", err)
		}

		defer f.Close()

		_, err = toml.NewDecoder(f).Decode(&cfg)
		if err != nil {
			return cfg, fmt.Errorf("decode toml: %w", err)
		}
	case ".env":
		err := godotenv.Load(filename)
		if err != nil {
			return cfg, fmt.Errorf("load env: %w", err)
		}
	}

	if useEnv || ext == ".env" {
		err := envconfig.Process("APP", &cfg)
		if err != nil {
			return cfg, fmt.Errorf("decode env: %w", err)
		}
	}

	return cfg, nil
}
