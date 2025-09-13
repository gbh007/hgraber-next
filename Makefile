TAG = $(shell git tag -l --points-at HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

MOD_NAME = github.com/gbh007/hgraber-next
LDFLAGS = -ldflags "-X '$(MOD_NAME)/version.Version=$(TAG)' -X '$(MOD_NAME)/version.Commit=$(COMMIT)' -X '$(MOD_NAME)/version.BuildAt=$(BUILD_TIME)' -X '$(MOD_NAME)/version.Branch=$(BRANCH)'"
SERVICE_BIN = $(PWD)/bin/build

GOBIN = $(PWD)/bin/utils
GOLANGCI_LINT = $(GOBIN)/golangci-lint

$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(GOBIN) v2.4.0

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run

.PHONY: install-tools
install-tools:
# 	На данный момент не работает корректно
# 	go get -u -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0
	go get -u -tool github.com/ogen-go/ogen/cmd/ogen@v1.14.0
	go get -u -tool golang.org/x/tools/cmd/deadcode@v0.36.0
	go get -u -tool mvdan.cc/gofumpt@v0.8.0
	go get -u -tool golang.org/x/tools/cmd/goimports@v0.36.0
	go get -u -tool github.com/daixiang0/gci@v0.13.7

# .PHONY: lint
# lint:
# 	go tool golangci-lint run

.PHONY: deadcode
deadcode:
	go tool deadcode -test ./...

.PHONY: format
format:
	go tool gofumpt -l -w .
	go tool goimports -l -w .
	go tool gci write -s standard -s default -s "prefix(github.com/gbh007/hgraber-next)" --skip-generated .

.PHONY: generate
generate:
	go tool ogen --target openapi/agentapi -package agentapi --clean openapi/agent.yaml
	go tool ogen --target openapi/serverapi -package serverapi --clean openapi/server.yaml

.PHONY: tidy
tidy:
	go mod tidy

create_build_dir:
	mkdir -p $(SERVICE_BIN)

.PHONY: run-example
run-example: create_build_dir
	go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/server  ./cmd/server

	$(SERVICE_BIN)/server --config config-example.toml
 
.PHONY: run
run: create_build_dir
	go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/server  ./cmd/server

	$(SERVICE_BIN)/server --config config.toml


.PHONY: config
config: create_build_dir
	go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/configremaper  ./cmd/configremaper

	$(SERVICE_BIN)/configremaper --out config-generated.yaml
	$(SERVICE_BIN)/configremaper --out config-generated.env
	$(SERVICE_BIN)/configremaper --out config-generated.toml

.PHONY: grafana
grafana: create_build_dir
	go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/grafanagenerator  ./cmd/grafanagenerator

	$(SERVICE_BIN)/grafanagenerator --config config-dashboard.toml


.PHONY: build
build: create_build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/server-linux-arm64  ./cmd/server
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/server-linux-amd64  ./cmd/server
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/server-windows-amd64.exe  ./cmd/server

.PHONY: jsonnet-build
jsonnet-build:
	cd jsonnet && make build

.PHONY: jsonnet-custom
jsonnet-custom:
	cd jsonnet && make custom