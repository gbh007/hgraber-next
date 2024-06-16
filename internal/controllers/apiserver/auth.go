package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
)

// FIXME: реализовать
func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t server.HeaderAuth) (context.Context, error) {
	return ctx, nil
}
