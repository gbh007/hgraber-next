#! /bin/bash

TAG=$(git tag -l --points-at HEAD)
COMMIT=$(git rev-parse --short HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

MOD_NAME="github.com/gbh007/hgraber-next"
LDFLAGS="-X '${MOD_NAME}/version.Version=${TAG}' -X '${MOD_NAME}/version.Commit=${COMMIT}' -X '${MOD_NAME}/version.BuildAt=${BUILD_TIME}' -X '${MOD_NAME}/version.Branch=${BRANCH}'"

go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/server  ./cmd/server

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/server-linux-arm64  ./cmd/server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/server-linux-amd64  ./cmd/server
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/server-windows-amd64.exe  ./cmd/server