package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

//nolint:cyclop,funlen // будет исправлено позднее
func ImportConfig[T any](filename string, newConfig func() T) (_ T, returnedErr error) {
	var cfg T

	if newConfig != nil {
		cfg = newConfig()
	}

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

	return cfg, nil
}
