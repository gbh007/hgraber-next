package agent

import (
	"flag"
	"log/slog"
	"os"

	"github.com/gbh007/hgraber-next/external/config"
	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
)

func parseConfig() (agnconfig.AgentConfig, bool, error) {
	configPath := flag.String("config", "config.toml", "path to config")
	scan := flag.Bool("scan", false, "scan zip file to register in db")

	flag.Parse()

	c, err := config.ImportConfig(*configPath, agnconfig.AgentDefaultConfigWrapped())

	return c, *scan, err
}

func initLogger(cfg agnconfig.AgentConfig) *slog.Logger {
	slogOpt := &slog.HandlerOptions{
		AddSource: cfg.Log.IncludeSource,
		Level:     cfg.Log.SlogLevel(),
	}

	return slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			slogOpt,
		),
	)
}
