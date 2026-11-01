package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

func ExportToFile[T any](cfg T, filename string) (returnedErr error) {
	f, err := os.Create(filename) //nolint:gosec // путь может быть любым доступным
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() {
		err := f.Close()

		switch {
		case returnedErr != nil && err != nil:
			returnedErr = errors.Join(returnedErr, fmt.Errorf("close file: %w", err))
		case err != nil:
			returnedErr = fmt.Errorf("close file: %w", err)
		default:
		}
	}()

	err = ExportToWriter(f, cfg)
	if err != nil {
		return fmt.Errorf("export to writer: %w", err)
	}

	return nil
}

func ExportToWriter[T any](w io.Writer, cfg T) error {
	enc := toml.NewEncoder(w)
	enc.Indent = ""

	err := enc.Encode(cfg)
	if err != nil {
		return fmt.Errorf("encode toml: %w", err)
	}

	return nil
}
