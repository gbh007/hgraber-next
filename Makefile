TAG = $(shell git tag -l --points-at HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

MOD_NAME = github.com/gbh007/hgraber-next
LDFLAGS = -ldflags "-X '$(MOD_NAME)/version.Version=$(TAG)' -X '$(MOD_NAME)/version.Commit=$(COMMIT)' -X '$(MOD_NAME)/version.BuildAt=$(BUILD_TIME)' -X '$(MOD_NAME)/version.Branch=$(BRANCH)'"


OGEN = github.com/ogen-go/ogen/cmd/ogen@v1.8.1

.PHONY: generate
generate:
	go run $(OGEN) --target openapi/agentapi -package agentapi --clean openapi/agent.yaml
	go run $(OGEN) --target openapi/serverapi -package serverapi --clean openapi/server.yaml

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


.PHONY: build
build: create_build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o ./_build/server-linux-arm64  ./cmd/server
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o ./_build/server-linux-amd64  ./cmd/server
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o ./_build/server-windows-amd64.exe  ./cmd/server