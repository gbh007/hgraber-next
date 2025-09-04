package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/gbh007/hgraber-next/config"
)

func main() {
	inPath := flag.String("in", "", "path to input config")
	outPath := flag.String("out", "", "path to out config")
	useEnv := flag.Bool("env", false, "use environment")
	flag.Parse()

	logger := slog.Default()

	if len(*outPath) == 0 {
		logger.Error("empty out path")
		os.Exit(1)
	}

	cfg, err := config.ImportConfig(*inPath, *useEnv)
	if err != nil {
		logger.Error("import config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	err = config.ExportToFile(&cfg, *outPath)
	if err != nil {
		logger.Error("export config", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
