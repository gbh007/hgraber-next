package legacyadapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type Page struct {
	PageNumber int
	URL        string
	Ext        string
}

type legacyParser interface {
	Name(ctx context.Context) (string, error)
	Pages(ctx context.Context) ([]Page, error)
	Tags(ctx context.Context) ([]string, error)
	Authors(ctx context.Context) ([]string, error)
	Characters(ctx context.Context) ([]string, error)
	Languages(ctx context.Context) ([]string, error)
	Categories(ctx context.Context) ([]string, error)
	Parodies(ctx context.Context) ([]string, error)
	Groups(ctx context.Context) ([]string, error)
}

type Adapter struct {
	old legacyParser
	u   url.URL
}

func New(old legacyParser, u url.URL) Adapter {
	return Adapter{
		old: old,
		u:   u,
	}
}

func (Adapter) HProxyBookDetails(ctx context.Context, pageLimit *int) (parserdto.Book, error) {
	return parserdto.Book{}, nil
}

func (adapter Adapter) BookDetails(ctx context.Context) (serveragentdto.AgentBookDetails, error) {
	details := serveragentdto.AgentBookDetails{
		URL: adapter.u,
	}

	var err error

	details.Name, err = adapter.old.Name(ctx)
	if err != nil {
		return serveragentdto.AgentBookDetails{}, fmt.Errorf("name: %w", err)
	}

	pages, err := adapter.old.Pages(ctx)
	if err != nil {
		return serveragentdto.AgentBookDetails{}, fmt.Errorf("pages: %w", err)
	}

	details.PageCount = len(pages)
	details.Pages, err = pkg.MapWithError(pages, func(p Page) (serveragentdto.AgentBookDetailsPagesItem, error) {
		u, err := url.Parse(p.URL)
		if err != nil {
			return serveragentdto.AgentBookDetailsPagesItem{}, fmt.Errorf("page %d: %w", p.PageNumber, err)
		}

		return serveragentdto.AgentBookDetailsPagesItem{
			PageNumber: p.PageNumber,
			URL:        *u,
			Filename:   fmt.Sprintf("%d.%s", p.PageNumber, p.Ext),
		}, nil
	})
	if err != nil {
		return serveragentdto.AgentBookDetails{}, fmt.Errorf("convert pages: %w", err)
	}

	for _, attrCode := range core.AllAttributes {
		values, err := ParseBookAttr(ctx, adapter.old, attrCode)
		if err != nil {
			return serveragentdto.AgentBookDetails{}, fmt.Errorf("%s: %w", string(attrCode), err)
		}

		if len(values) > 0 {
			details.Attributes = append(details.Attributes, serveragentdto.AgentBookDetailsAttributesItem{
				Code:   string(attrCode),
				Values: values,
			})
		}
	}

	return details, nil
}
