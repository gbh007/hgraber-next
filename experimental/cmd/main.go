package main

import (
	"encoding/json"
	"flag"
	"net/url"
	"os"

	"github.com/go-openapi/strfmt"
	goapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/models"

	"github.com/gbh007/hgraber-next/config"
	"github.com/gbh007/hgraber-next/experimental/generator"
)

//nolint:revive // будет исправлено позднее
type Config struct {
	Grafana struct {
		Addr   string `toml:"addr"`
		Token  string `toml:"token"`
		Folder string `toml:"folder"`
		UID    string `toml:"uid"`
	} `toml:"grafana"`
	HGraber struct {
		Services []string `toml:"services"`
	} `toml:"hgraber"`
}

func main() {
	configPath := flag.String("config", "config.toml", "path to config")
	flag.Parse()

	cfg, err := config.ImportConfig(*configPath, true, func() Config { return Config{} })
	if err != nil {
		panic(err)
	}

	// TODO: подумать о необходимости
	// if cfg.Grafana.Folder == "" {
	// 	panic("empty folder")
	// }

	if cfg.Grafana.UID == "" {
		panic("empty uid")
	}

	u, err := url.Parse(cfg.Grafana.Addr)
	if err != nil {
		panic(err)
	}

	g := generator.New(cfg.Grafana.UID, cfg.HGraber.Services)

	transportCfg := &goapi.TransportConfig{
		Host:     u.Host,
		BasePath: u.Path,
		Schemes:  []string{u.Scheme},
		APIKey:   cfg.Grafana.Token,
	}

	client := goapi.NewHTTPClientWithConfig(strfmt.Default, transportCfg)

	dashboardModel, err := g.Build()
	if err != nil {
		panic(err)
	}

	response, err := client.Dashboards.PostDashboard(&models.SaveDashboardCommand{
		FolderUID: cfg.Grafana.Folder,
		Dashboard: dashboardModel,
		Overwrite: true,
	})
	if err != nil {
		panic(err)
	}

	if *response.Payload.Status != "success" {
		panic(*response.Payload.Status)
	}

	out, err := os.Create("grafana-dashboard (experimental).json")
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(out)
	enc.SetIndent("", "   ")

	err = enc.Encode(dashboardModel)
	if err != nil {
		panic(err)
	}

	err = out.Close()
	if err != nil {
		panic(err)
	}
}
