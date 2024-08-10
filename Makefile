TAG = $(shell git tag -l --points-at HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS = -ldflags "-X 'hgnext/internal/version.Version=$(TAG)' -X 'hgnext/internal/version.Commit=$(COMMIT)' -X 'hgnext/internal/version.BuildAt=$(BUILD_TIME)' -X 'hgnext/internal/version.Branch=$(BRANCH)'"


.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target open_api/agentAPI -package agentAPI --clean open_api/agent.yaml
	go run github.com/ogen-go/ogen/cmd/ogen --target open_api/serverAPI -package serverAPI --clean open_api/server.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: run
run: create_build_dir
	go build $(LDFLAGS) -trimpath -o ./_build/server  ./cmd/server

	./_build/server --config config-example.yaml

.PHONY: runafs
runafs: create_build_dir
	go build $(LDFLAGS) -trimpath -o ./_build/server  ./cmd/server

	APP_STORAGE_FS_AGENT_ID=019067fc-8d4f-769d-8c4f-e755597f9525 \
	APP_APPLICATION_TRACE_ENDPOINT=http://localhost:4318/v1/traces \
	./_build/server --config config-example.yaml