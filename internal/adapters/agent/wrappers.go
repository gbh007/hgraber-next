package agent

import (
	"context"
	"io"
	"net/url"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (c *Client) BookParse(ctx context.Context, agentID uuid.UUID, url url.URL) (entities.AgentBookDetails, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return entities.AgentBookDetails{}, err
	}

	return adapter.BookParse(ctx, url)
}

func (c *Client) BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.BooksCheck(ctx, urls)
}

func (c *Client) ExportArchive(ctx context.Context, agentID uuid.UUID, bookID uuid.UUID, bookName string, body io.Reader) error {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return err
	}

	return adapter.ExportArchive(ctx, bookID, bookName, body)
}

func (c *Client) PageLoad(ctx context.Context, agentID uuid.UUID, url entities.AgentPageURL) (io.Reader, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.PageLoad(ctx, url)
}

func (c *Client) PagesCheck(ctx context.Context, agentID uuid.UUID, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.PagesCheck(ctx, urls)
}

func (c *Client) Status(ctx context.Context, agentID uuid.UUID) (entities.AgentStatus, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return entities.AgentStatus{}, err
	}

	return adapter.Status(ctx)
}
