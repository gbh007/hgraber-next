package agentapi

import (
	"context"
	"io"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

func (c *Client) BookParse(ctx context.Context, agentID uuid.UUID, u url.URL) (serveragentdto.AgentBookDetails, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return serveragentdto.AgentBookDetails{}, err
	}

	return adapter.BookParse(ctx, u) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) BooksCheck(
	ctx context.Context,
	agentID uuid.UUID,
	urls []url.URL,
) ([]serveragentdto.AgentBookCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.BooksCheck(ctx, urls) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) BooksCheckMultiple(
	ctx context.Context,
	agentID uuid.UUID,
	u url.URL,
) ([]serveragentdto.AgentBookCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.BooksCheckMulti(ctx, u) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) ExportArchive(ctx context.Context, agentID uuid.UUID, data serveragentdto.AgentExportData) error {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return err
	}

	return adapter.ExportArchive(ctx, data) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) PageLoad(ctx context.Context, agentID uuid.UUID, u serveragentdto.AgentPageURL) (io.Reader, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.PageLoad(ctx, u) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) PagesCheck(
	ctx context.Context,
	agentID uuid.UUID,
	urls []serveragentdto.AgentPageURL,
) ([]serveragentdto.AgentPageCheckResult, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return nil, err
	}

	return adapter.PagesCheck(ctx, urls) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) Status(ctx context.Context, agentID uuid.UUID) (serveragentdto.AgentStatus, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return serveragentdto.AgentStatus{}, err
	}

	return adapter.Status(ctx) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) HProxyList(ctx context.Context, agentID uuid.UUID, u url.URL) (hproxymodel.List, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return hproxymodel.List{}, err
	}

	return adapter.HProxyList(ctx, u) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}

func (c *Client) HProxyBook(
	ctx context.Context,
	agentID uuid.UUID,
	u url.URL,
	pageLimit *int,
) (hproxymodel.Book, error) {
	adapter, err := c.getAdapter(agentID)
	if err != nil {
		return hproxymodel.Book{}, err
	}

	return adapter.HProxyBook(ctx, u, pageLimit) //nolint:wrapcheck // это обычный врапер, обертка не требуется
}
