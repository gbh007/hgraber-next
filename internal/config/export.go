package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

const envTemplate = `{{ range . }}{{ .Key }}={{ .Field }}
{{ end }}`

func ExportToFile[T any](cfg *T, filename string) (returnedErr error) {
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

	ext := path.Ext(filename)

	err = ExportToWriter(f, cfg, ext)
	if err != nil {
		return fmt.Errorf("export to writer: %w", err)
	}

	return nil
}

func ExportToWriter[T any](w io.Writer, cfg *T, ext string) error {
	switch ext {
	case ConfigExtYml, ConfigExtYaml:
		enc := yaml.NewEncoder(w)
		enc.SetIndent(YamlIndent)

		err := enc.Encode(cfg)
		if err != nil {
			return fmt.Errorf("encode yaml: %w", err)
		}

	case ConfigExtToml:
		enc := toml.NewEncoder(w)
		enc.Indent = ""

		err := enc.Encode(cfg)
		if err != nil {
			return fmt.Errorf("encode toml: %w", err)
		}

	case ConfigExtEnv:
		err := envconfig.Usaget("APP", cfg, w, template.Must(template.New("cfg").Parse(envTemplate)))
		if err != nil {
			return fmt.Errorf("encode env usage: %w", err)
		}
	}

	return nil
}
