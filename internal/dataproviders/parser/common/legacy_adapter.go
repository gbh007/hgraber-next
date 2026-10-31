package common

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type legacyParser interface {
	Name(ctx context.Context) (string, error)
	Pages(ctx context.Context) ([]parsing.RawPage, error)
	Tags(ctx context.Context) ([]string, error)
	Authors(ctx context.Context) ([]string, error)
	Characters(ctx context.Context) ([]string, error)
	Languages(ctx context.Context) ([]string, error)
	Categories(ctx context.Context) ([]string, error)
	Parodies(ctx context.Context) ([]string, error)
	Groups(ctx context.Context) ([]string, error)
}

func ParseBookAttr(ctx context.Context, p legacyParser, attr string) ([]string, error) {
	switch attr {
	case core.AttributeCodeAuthor:
		return p.Authors(ctx)

	case core.AttributeCodeCategory:
		return p.Categories(ctx)

	case core.AttributeCodeCharacter:
		return p.Characters(ctx)

	case core.AttributeCodeGroup:
		return p.Groups(ctx)

	case core.AttributeCodeLanguage:
		return p.Languages(ctx)

	case core.AttributeCodeParody:
		return p.Parodies(ctx)

	case core.AttributeCodeTag:
		return p.Tags(ctx)

	default:
		return []string{}, parsing.UnknownAttributeError
	}
}

type LegacyParserAdapter struct {
	BookParser
	old legacyParser
	u   url.URL
}

func NewLegacyParserAdapter(old legacyParser, u url.URL) LegacyParserAdapter {
	return LegacyParserAdapter{
		old: old,
		u:   u,
	}
}

func (adapter LegacyParserAdapter) BookDetails(ctx context.Context) (serveragentdto.AgentBookDetails, error) {
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
	details.Pages, err = pkg.MapWithError(pages, func(p parsing.RawPage) (serveragentdto.AgentBookDetailsPagesItem, error) {
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
