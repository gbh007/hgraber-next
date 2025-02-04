package agent

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
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

func (c *Client) FSState(ctx context.Context, agentID uuid.UUID, includeFileIDs, includeFileSizes bool) (core.FSState, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return core.FSState{}, err
	}

	return adapter.ToFS().State(ctx, includeFileIDs, includeFileSizes)
}

func (c *Client) CreateHighwayToken(ctx context.Context, agentID uuid.UUID) (string, time.Time, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return "", time.Time{}, err
	}

	return adapter.ToFS().CreateHighwayToken(ctx)
}
