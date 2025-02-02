#! /bin/bash

TAG=$(git tag -l --points-at HEAD)
COMMIT=$(git rev-parse --short HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

MOD_NAME="github.com/gbh007/hgraber-next"
LDFLAGS="-X '${MOD_NAME}/internal/version.Version=${TAG}' -X '${MOD_NAME}/internal/version.Commit=${COMMIT}' -X '${MOD_NAME}/internal/version.BuildAt=${BUILD_TIME}' -X '${MOD_NAME}/internal/version.Branch=${BRANCH}'"

go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/server  ./cmd/server
