.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/adapters/agent/internal/client -package client --clean open_api_specs/agent.yaml
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/controllers/apiserver/internal/server -package server --clean open_api_specs/server.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: run
run: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	./_build/server --config config-example.yaml

.PHONY: runafs
runafs: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	APP_STORAGE_FS_AGENT_ID=019067fc-8d4f-769d-8c4f-e755597f9525 \
	OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318 \
	APP_WORKERS_PAGE_COUNT=1 \
	./_build/server --config config-example.yaml