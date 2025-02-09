JSONNET=github.com/google/go-jsonnet/cmd/jsonnet@latest
JSONNET_JB=github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb@latest

.PHONY: jsonnet
jsonnet:
	go run $(JSONNET_JB) install github.com/grafana/grafonnet/gen/grafonnet-latest@main
	go run $(JSONNET) -J vendor dashboard.jsonnet > dashboard.json