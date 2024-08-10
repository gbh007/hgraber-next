package adapter

import (
	"context"

	"hgnext/open_api/agentAPI"
)

type securitySource struct {
	token string
}

func (s securitySource) HeaderAuth(ctx context.Context, operationName string) (agentAPI.HeaderAuth, error) {
	return agentAPI.HeaderAuth{
		APIKey: s.token,
	}, nil
}
