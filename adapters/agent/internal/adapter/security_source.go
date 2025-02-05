package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

type securitySource struct {
	token string
}

func (s securitySource) HeaderAuth(ctx context.Context, operationName string) (agentapi.HeaderAuth, error) {
	return agentapi.HeaderAuth{
		APIKey: s.token,
	}, nil
}
