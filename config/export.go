package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func ExportToFile[T any](cfg *T, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer f.Close()

	ext := path.Ext(filename)

	err = ExportToWriter(f, cfg, ext)
	if err != nil {
		return fmt.Errorf("export to writer: %w", err)
	}

	return nil
}

func ExportToWriter[T any](w io.Writer, cfg *T, ext string) error {
	switch ext {
	case ".yml", ".yaml":
		enc := yaml.NewEncoder(w)
		enc.SetIndent(2)

		err := enc.Encode(cfg)
		if err != nil {
			return fmt.Errorf("encode yaml: %w", err)
		}
	case ".toml":
		enc := toml.NewEncoder(w)
		enc.Indent = ""

		err := enc.Encode(cfg)
		if err != nil {
			return fmt.Errorf("encode toml: %w", err)
		}
	case ".env":
		err := envconfig.Usaget("APP", cfg, w, template.Must(template.New("cfg").Parse(envTemplate)))
		if err != nil {
			return fmt.Errorf("encode env usage: %w", err)
		}
	}

	return nil
}

const envTemplate = `{{ range . }}{{ .Key }}={{ .Field }}
{{ end }}`
