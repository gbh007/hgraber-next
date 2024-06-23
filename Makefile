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
	./_build/server