TAG = $(shell git tag -l --points-at HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

MOD_NAME = github.com/gbh007/hgraber-next
LDFLAGS = -ldflags "-X '$(MOD_NAME)/internal/version.Version=$(TAG)' -X '$(MOD_NAME)/internal/version.Commit=$(COMMIT)' -X '$(MOD_NAME)/internal/version.BuildAt=$(BUILD_TIME)' -X '$(MOD_NAME)/internal/version.Branch=$(BRANCH)'"


OGEN = github.com/ogen-go/ogen/cmd/ogen@v1.8.1

.PHONY: generate
generate:
	go run $(OGEN) --target open_api/agentAPI -package agentAPI --clean open_api/agent.yaml
	go run $(OGEN) --target open_api/serverAPI -package serverAPI --clean open_api/server.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: run-example
run-example: create_build_dir
	go build $(LDFLAGS) -trimpath -o ./_build/server  ./cmd/server

	./_build/server --config config-example.yaml
 
.PHONY: run
run: create_build_dir
	go build $(LDFLAGS) -trimpath -o ./_build/server  ./cmd/server

	./_build/server


.PHONY: config
config: create_build_dir
	go build $(LDFLAGS) -trimpath -o ./_build/server  ./cmd/server

	./_build/server --generate-config config-generated.yaml