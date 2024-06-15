package adapter

import (
	"context"

	"hgnext/internal/adapters/agent/internal/client"
)

type securitySource struct {
	token string
}

func (s securitySource) HeaderAuth(ctx context.Context, operationName string) (client.HeaderAuth, error) {
	return client.HeaderAuth{
		APIKey: s.token,
	}, nil
}
