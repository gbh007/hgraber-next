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

func ImportConfig(filename string, useEnv bool) (Config, error) {
	c := ConfigDefault()

	ext := path.Ext(filename)

	switch ext {
	case ".yml", ".yaml":
		f, err := os.Open(filename)
		if err != nil {
			return Config{}, fmt.Errorf("open config file: %w", err)
		}

		defer f.Close()

		err = yaml.NewDecoder(f).Decode(&c)
		if err != nil {
			return Config{}, fmt.Errorf("decode yaml: %w", err)
		}
	case ".toml":
		f, err := os.Open(filename)
		if err != nil {
			return Config{}, fmt.Errorf("open config file: %w", err)
		}

		defer f.Close()

		_, err = toml.NewDecoder(f).Decode(&c)
		if err != nil {
			return Config{}, fmt.Errorf("decode toml: %w", err)
		}
	case ".env":
		err := godotenv.Load(filename)
		if err != nil {
			return Config{}, fmt.Errorf("load env: %w", err)
		}
	}

	if useEnv || ext == ".env" {
		err := envconfig.Process("APP", &c)
		if err != nil {
			return Config{}, fmt.Errorf("decode env: %w", err)
		}
	}

	return c, nil
}
