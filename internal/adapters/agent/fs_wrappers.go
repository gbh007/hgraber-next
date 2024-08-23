package agent

import (
	"context"
	"io"

	"github.com/google/uuid"
)

func (c *Client) FSCreate(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID, body io.Reader) error {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return err
	}

	return adapter.ToFS().Create(ctx, fileID, body)
}

func (c *Client) FSDelete(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) error {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return err
	}

	return adapter.ToFS().Delete(ctx, fileID)
}

func (c *Client) FSGet(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) (io.Reader, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.ToFS().Get(ctx, fileID)
}

func (c *Client) FSIDs(ctx context.Context, agentID uuid.UUID) ([]uuid.UUID, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.ToFS().IDs(ctx)
}
