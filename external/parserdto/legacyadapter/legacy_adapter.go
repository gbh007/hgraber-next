package legacyadapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/samber/lo"
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

func (adapter Adapter) Book(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error) {
	return Convert(ctx, adapter.u, adapter.old)
}

func Convert(ctx context.Context, u url.URL, la legacyParser) (parserdto.Book, error) {
	details := parserdto.Book{
		URL: u,
	}

	var err error

	details.Name, err = la.Name(ctx)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("name: %w", err)
	}

	pages, err := la.Pages(ctx)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("pages: %w", err)
	}

	details.PageCount = len(pages)

	details.Pages, err = lo.MapErr(pages, func(p Page, _ int) (parserdto.Page, error) {
		u, err := url.Parse(p.URL)
		if err != nil {
			return parserdto.Page{}, fmt.Errorf("page %d: %w", p.PageNumber, err)
		}

		return parserdto.Page{
			PageNumber: p.PageNumber,
			URL:        *u,
			Filename:   fmt.Sprintf("%d.%s", p.PageNumber, p.Ext),
		}, nil
	})
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("convert pages: %w", err)
	}

	for _, attrCode := range core.AllAttributes {
		values, err := ParseBookAttr(ctx, la, attrCode)
		if err != nil {
			return parserdto.Book{}, fmt.Errorf("%s: %w", attrCode, err)
		}

		if len(values) > 0 {
			details.Attributes = append(details.Attributes, parserdto.Attribute{
				Code: attrCode,
				Values: lo.Map(values, func(v string, _ int) parserdto.AttributeValue {
					return parserdto.AttributeValue{
						Value: v,
					}
				}),
			})
		}
	}

	return details, nil
}
