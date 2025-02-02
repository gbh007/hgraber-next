package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

type securitySource struct {
	token string
}

func (s securitySource) HeaderAuth(ctx context.Context, operationName string) (agentAPI.HeaderAuth, error) {
	return agentAPI.HeaderAuth{
		APIKey: s.token,
	}, nil
}
