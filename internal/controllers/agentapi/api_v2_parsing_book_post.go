package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/samber/lo"
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
		Pages: lo.Map(details.Pages, func(p parserdto.Page, _ int) agentapi.BookDetailsPagesItem {
			return agentapi.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
		Attributes: lo.Map(details.Attributes, func(attr parserdto.Attribute, _ int) agentapi.BookDetailsAttributesItem {
			return agentapi.BookDetailsAttributesItem{
				Code: attr.Code,
				Values: lo.Map(attr.Values, func(v parserdto.AttributeValue, _ int) agentapi.BookDetailsAttributesItemValuesItem {
					return agentapi.BookDetailsAttributesItemValuesItem{
						Name: v.Value,
						URL:  OptURL(v.URL),
					}
				}),
			}
		}),
	}, nil
}
