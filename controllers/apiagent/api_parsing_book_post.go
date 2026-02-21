package apiagent

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingBookPost(
	ctx context.Context,
	req *agentapi.APIParsingBookPostReq,
) (*agentapi.BookDetails, error) {
	details, err := c.parsingUseCases.BookByURL(ctx, req.URL)
	if err != nil {
		return nil, apiError{
			Code:      http.StatusInternalServerError,
			InnerCode: ParseUseCaseCode,
			Details:   err.Error(),
		}
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
					Code:   agentapi.BookDetailsAttributesItemCode(code),
					Values: values,
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
