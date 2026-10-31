package agentcacheapi

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/samber/lo"
)

func (c *Controller) APIV2ParsingBookPost(
	ctx context.Context,
	req *agentapi.APIV2ParsingBookPostReq,
) (*agentapi.BookDetails, error) {
	details, err := c.parsingUseCases.BookByURL(ctx, req.URL)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	var u url.URL

	if details.Book.OriginURL != nil {
		u = *details.Book.OriginURL
	}

	return &agentapi.BookDetails{
		URL:       u,
		Name:      details.Book.Name,
		PageCount: details.Book.PageCount,
		Attributes: pkg.MapToSlice(
			details.Attributes,
			func(code string, values []string) agentapi.BookDetailsAttributesItem {
				return agentapi.BookDetailsAttributesItem{
					Code: code,
					Values: lo.Map(values, func(v string, _ int) agentapi.BookDetailsAttributesItemValuesItem {
						return agentapi.BookDetailsAttributesItemValuesItem{
							Name: v,
						}
					}),
				}
			},
		),
		Pages: pkg.Map(details.Pages, func(p core.Page) agentapi.BookDetailsPagesItem {
			var u url.URL

			if p.OriginURL != nil {
				u = *p.OriginURL
			}

			return agentapi.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        u,
				Filename:   p.Filename(),
			}
		}),
	}, nil
}
