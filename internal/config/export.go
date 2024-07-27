package config

import (
	"fmt"
	"os"
	"text/template"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func ExportToFile(cfg Config, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(2)

	err = enc.Encode(cfg)
	if err != nil {
		return fmt.Errorf("encode yaml: %w", err)
	}

	err = envconfig.Usaget("APP", &cfg, f, template.Must(template.New("cfg").Parse(envTemplate)))
	if err != nil {
		return fmt.Errorf("encode env usage: %w", err)
	}

	return nil
}

const envTemplate = `
{{ range . }}# {{ .Key }}={{ .Field }}
{{ end }}`
