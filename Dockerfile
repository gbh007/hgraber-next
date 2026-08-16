FROM node:22-alpine AS node-builder

WORKDIR /build

COPY frontend/package*.json ./
RUN --mount=type=cache,target="/build/node_modules" npm install

COPY --exclude=frontend/dist --exclude=frontend/node_modules frontend/ ./
RUN --mount=type=cache,target="/build/node_modules" npm run build

FROM golang:1.26.5 AS go-builder

# TODO: подумать над включением
# RUN useradd -u 1000 nonroot
WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target="$(go env GOMODCACHE)" \
    --mount=type=cache,target="$(go env GOCACHE)" \
    go mod download

COPY --exclude=frontend . .
RUN --mount=type=cache,target="$(go env GOMODCACHE)" \
    --mount=type=cache,target="$(go env GOCACHE)" \
    CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /build/main \
    ./cmd/server

RUN --mount=type=cache,target="$(go env GOMODCACHE)" \
    --mount=type=cache,target="$(go env GOCACHE)" \
    CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /build/grafanagenerator \
    ./cmd/grafanagenerator

FROM scratch

COPY --from=go-builder /etc/passwd /etc/passwd
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=go-builder /build/main /app/main
COPY --from=go-builder /build/grafanagenerator /app/grafanagenerator
COPY --from=node-builder /build/dist /app/static

ENV APP_API_STATIC_DIR=/app/static

# TODO: подумать над включением
# USER nonroot
WORKDIR /app

ENTRYPOINT ["/app/main"]
