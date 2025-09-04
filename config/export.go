package config

import (
	"fmt"
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

	switch path.Ext(filename) {
	case ".yml", ".yaml":
		enc := yaml.NewEncoder(f)
		enc.SetIndent(2)

		err = enc.Encode(cfg)
		if err != nil {
			return fmt.Errorf("encode yaml: %w", err)
		}
	case ".toml":
		enc := toml.NewEncoder(f)
		enc.Indent = ""

		err = enc.Encode(cfg)
		if err != nil {
			return fmt.Errorf("encode toml: %w", err)
		}
	case ".env":
		err = envconfig.Usaget("APP", cfg, f, template.Must(template.New("cfg").Parse(envTemplate)))
		if err != nil {
			return fmt.Errorf("encode env usage: %w", err)
		}
	}

	return nil
}

const envTemplate = `{{ range . }}{{ .Key }}={{ .Field }}
{{ end }}`
