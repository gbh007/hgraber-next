.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/adapters/agent/internal/client -package client --clean open_api_specs/agent.yaml
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/controllers/apiserver/internal/server -package server --clean open_api_specs/server.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: run
run: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	APP_POSTGRESQL_CONNECTION=postgres://hgrabernextuser:hgrabernextpass@localhost:5432/hgrabernext?sslmode=disable \
	APP_FILE_PATH=./.hidden/filedata \
	APP_WEB_SERVER_ADDR=127.0.0.1:8080 \
	APP_EXTERNAL_WEB_SERVER_ADDR=http://localhost:8080 \
	APP_DEBUG=true \
	APP_WEB_STATIC_DIR=internal/controllers/apiserver/internal/static \
	APP_API_TOKEN=local-next \
	APP_METRIC_TIMEOUT=100ms \
	./_build/server

.PHONY: runafs
runafs: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	APP_POSTGRESQL_CONNECTION=postgres://hgrabernextuser:hgrabernextpass@localhost:5432/hgrabernext?sslmode=disable \
	APP_WEB_SERVER_ADDR=127.0.0.1:8080 \
	APP_EXTERNAL_WEB_SERVER_ADDR=http://localhost:8080 \
	APP_DEBUG=true \
	APP_WEB_STATIC_DIR=internal/controllers/apiserver/internal/static \
	APP_API_TOKEN=local-next \
	APP_METRIC_TIMEOUT=100ms \
	APP_FS_AGENT_ID=019067fc-8d4f-769d-8c4f-e755597f9525 \
	./_build/server