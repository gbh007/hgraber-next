package main

import (
	"flag"

	"hgnext/internal/config"
)

func main() {
	filename := flag.String("f", "", "filename to export config")
	flag.Parse()

	if *filename == "" {
		return
	}

	err := config.ExportToFile(config.ConfigDefault(), *filename)
	if err != nil {
		panic(err)
	}
}
