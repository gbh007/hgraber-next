package agent

import (
	"context"
	"io"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
)

func (c *Client) BookParse(ctx context.Context, agentID uuid.UUID, url url.URL) (agentmodel.AgentBookDetails, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return agentmodel.AgentBookDetails{}, err
	}

	return adapter.BookParse(ctx, url)
}

func (c *Client) BooksCheck(
	ctx context.Context,
	agentID uuid.UUID,
	urls []url.URL,
) ([]agentmodel.AgentBookCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.BooksCheck(ctx, urls)
}

func (c *Client) BooksCheckMultiple(
	ctx context.Context,
	agentID uuid.UUID,
	u url.URL,
) ([]agentmodel.AgentBookCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.BooksCheckMulti(ctx, u)
}

func (c *Client) ExportArchive(ctx context.Context, agentID uuid.UUID, data agentmodel.AgentExportData) error {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return err
	}

	return adapter.ExportArchive(ctx, data)
}

func (c *Client) PageLoad(ctx context.Context, agentID uuid.UUID, url agentmodel.AgentPageURL) (io.Reader, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.PageLoad(ctx, url)
}

func (c *Client) PagesCheck(
	ctx context.Context,
	agentID uuid.UUID,
	urls []agentmodel.AgentPageURL,
) ([]agentmodel.AgentPageCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.PagesCheck(ctx, urls)
}

func (c *Client) Status(ctx context.Context, agentID uuid.UUID) (agentmodel.AgentStatus, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return agentmodel.AgentStatus{}, err
	}

	return adapter.Status(ctx)
}

func (c *Client) HProxyList(ctx context.Context, agentID uuid.UUID, u url.URL) (hproxymodel.List, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return hproxymodel.List{}, err
	}

	return adapter.HProxyList(ctx, u)
}

func (c *Client) HProxyBook(ctx context.Context, agentID uuid.UUID, u url.URL, pageLimit *int) (hproxymodel.Book, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return hproxymodel.Book{}, err
	}

	return adapter.HProxyBook(ctx, u, pageLimit)
}
