package common

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/hgraber"
	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type legacyParser interface {
	Name(ctx context.Context) (string, error)
	Pages(ctx context.Context) ([]hgraber.Page, error)
	Tags(ctx context.Context) ([]string, error)
	Authors(ctx context.Context) ([]string, error)
	Characters(ctx context.Context) ([]string, error)
	Languages(ctx context.Context) ([]string, error)
	Categories(ctx context.Context) ([]string, error)
	Parodies(ctx context.Context) ([]string, error)
	Groups(ctx context.Context) ([]string, error)
}

func ParseBookAttr(ctx context.Context, p legacyParser, attr hgraber.Attribute) ([]string, error) {
	switch attr {
	case hgraber.AttrAuthor:
		return p.Authors(ctx)

	case hgraber.AttrCategory:
		return p.Categories(ctx)

	case hgraber.AttrCharacter:
		return p.Characters(ctx)

	case hgraber.AttrGroup:
		return p.Groups(ctx)

	case hgraber.AttrLanguage:
		return p.Languages(ctx)

	case hgraber.AttrParody:
		return p.Parodies(ctx)

	case hgraber.AttrTag:
		return p.Tags(ctx)

	default:
		return []string{}, hgraber.UnknownAttributeError
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

func (adapter LegacyParserAdapter) BookDetails(ctx context.Context) (entities.AgentBookDetails, error) {
	details := entities.AgentBookDetails{
		URL: adapter.u,
	}

	var err error

	details.Name, err = adapter.old.Name(ctx)
	if err != nil {
		return entities.AgentBookDetails{}, fmt.Errorf("name: %w", err)
	}

	pages, err := adapter.old.Pages(ctx)
	if err != nil {
		return entities.AgentBookDetails{}, fmt.Errorf("pages: %w", err)
	}

	details.PageCount = len(pages)
	details.Pages, err = pkg.MapWithError(pages, func(p hgraber.Page) (entities.AgentBookDetailsPagesItem, error) {
		u, err := url.Parse(p.URL)
		if err != nil {
			return entities.AgentBookDetailsPagesItem{}, fmt.Errorf("page %d: %w", p.PageNumber, err)
		}

		return entities.AgentBookDetailsPagesItem{
			PageNumber: p.PageNumber,
			URL:        *u,
			Filename:   fmt.Sprintf("%d.%s", p.PageNumber, p.Ext),
		}, nil
	})
	if err != nil {
		return entities.AgentBookDetails{}, fmt.Errorf("convert pages: %w", err)
	}

	for _, attrCode := range hgraber.AllAttributes {
		values, err := ParseBookAttr(ctx, adapter.old, attrCode)
		if err != nil {
			return entities.AgentBookDetails{}, fmt.Errorf("%s: %w", string(attrCode), err)
		}

		if len(values) > 0 {
			details.Attributes = append(details.Attributes, entities.AgentBookDetailsAttributesItem{
				Code:   string(attrCode),
				Values: values,
			})
		}
	}

	return details, nil
}
