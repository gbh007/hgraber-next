package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2ParsingBookPost(
	ctx context.Context,
	req *agentapi.APIV2ParsingBookPostReq,
) (*agentapi.BookDetails, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	var pageLimit *int

	if req.PageLimit.Set {
		pageLimit = &req.PageLimit.Value
	}

	details, err := c.parsingUseCases.Book(ctx, req.URL, pageLimit)
	if err != nil {
		return nil, err
	}

	return &agentapi.BookDetails{
		Name:       details.Name,
		URL:        details.URL,
		PreviewURL: OptURL(details.PreviewURL),
		PageCount:  details.PageCount,
		Pages: pkg.Map(details.Pages, func(p parserdto.Page) agentapi.BookDetailsPagesItem {
			return agentapi.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
		Attributes: pkg.Map(details.Attributes, func(attr parserdto.Attribute) agentapi.BookDetailsAttributesItem {
			return agentapi.BookDetailsAttributesItem{
				Code: attr.Code,
				Values: pkg.Map(attr.Values, func(v parserdto.AttributeValue) agentapi.BookDetailsAttributesItemValuesItem {
					return agentapi.BookDetailsAttributesItemValuesItem{
						Name: v.Value,
						URL:  OptURL(v.URL),
					}
				}),
			}
		}),
	}, nil
}
