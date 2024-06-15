.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/adapters/agent/internal/client -package client --clean open_api_specs/agent.yaml
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/controllers/apiserver/internal/server -package server --clean open_api_specs/server.yaml