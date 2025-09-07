package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

//nolint:cyclop,funlen // будет исправлено позднее
func ImportConfig[T any](filename string, useEnv bool, newConfig func() T) (_ T, returnedErr error) {
	var cfg T

	if newConfig != nil {
		cfg = newConfig()
	}

	ext := path.Ext(filename)

	switch ext {
	case ConfigExtYml, ConfigExtYaml:
		f, err := os.Open(filename) //nolint:gosec // путь может быть любым доступным
		if err != nil {
			return cfg, fmt.Errorf("open config file: %w", err)
		}

		defer func() {
			err := f.Close()

			switch {
			case returnedErr != nil && err != nil:
				returnedErr = errors.Join(returnedErr, fmt.Errorf("close config file: %w", err))
			case err != nil:
				returnedErr = fmt.Errorf("close config file: %w", err)
			default:
			}
		}()

		err = yaml.NewDecoder(f).Decode(&cfg)
		if err != nil {
			return cfg, fmt.Errorf("decode yaml: %w", err)
		}
	case ConfigExtToml:
		f, err := os.Open(filename) //nolint:gosec // путь может быть любым доступным
		if err != nil {
			return cfg, fmt.Errorf("open config file: %w", err)
		}

		defer func() {
			err := f.Close()

			switch {
			case returnedErr != nil && err != nil:
				returnedErr = errors.Join(returnedErr, fmt.Errorf("close config file: %w", err))
			case err != nil:
				returnedErr = fmt.Errorf("close config file: %w", err)
			default:
			}
		}()

		_, err = toml.NewDecoder(f).Decode(&cfg)
		if err != nil {
			return cfg, fmt.Errorf("decode toml: %w", err)
		}
	case ConfigExtEnv:
		err := godotenv.Load(filename)
		if err != nil {
			return cfg, fmt.Errorf("load env: %w", err)
		}
	}

	if useEnv || ext == ConfigExtEnv {
		err := envconfig.Process("APP", &cfg)
		if err != nil {
			return cfg, fmt.Errorf("decode env: %w", err)
		}
	}

	return cfg, nil
}
