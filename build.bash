#! /bin/bash

TAG=$(git tag -l --points-at HEAD)
COMMIT=$(git rev-parse --short HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS="-X 'hgnext/internal/version.Version=${TAG}' -X 'hgnext/internal/version.Commit=${COMMIT}' -X 'hgnext/internal/version.BuildAt=${BUILD_TIME}' -X 'hgnext/internal/version.Branch=${BRANCH}'"

go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/server  ./cmd/server
