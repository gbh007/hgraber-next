JSONNET=github.com/google/go-jsonnet/cmd/jsonnet@latest
JSONNET_JB=github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb@latest

.PHONY: install
install:
	go run $(JSONNET_JB) install github.com/grafana/grafonnet/gen/grafonnet-latest@main

.PHONY: debug
debug: install
	go run $(JSONNET) -J vendor dashboard.jsonnet > dashboard.json

.PHONY: build
build: install
	go run $(JSONNET) -J vendor dashboard.jsonnet > ../grafana-dasboard.json