package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIHproxyParseBookPost(ctx context.Context, req *agentapi.APIHproxyParseBookPostReq) (*agentapi.APIHproxyParseBookPostOK, error) {
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

	details, err := c.parsingUseCases.HProxyBook(ctx, req.URL, pageLimit)
	if err != nil {
		return nil, err
	}

	return &agentapi.APIHproxyParseBookPostOK{
		Name:       details.Name,
		URL:        details.URL,
		PreviewURL: OptURL(details.PreviewURL),
		PageCount:  details.PageCount,
		Pages: pkg.Map(details.Pages, func(p agentdto.HProxyBookDetailsPagesItem) agentapi.APIHproxyParseBookPostOKPagesItem {
			return agentapi.APIHproxyParseBookPostOKPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
		Attributes: pkg.Map(details.Attributes, func(attr agentdto.HProxyBookDetailsAttributesItem) agentapi.APIHproxyParseBookPostOKAttributesItem {
			return agentapi.APIHproxyParseBookPostOKAttributesItem{
				Code: attr.Code,
				Values: pkg.Map(attr.Values, func(v agentdto.HProxyBookDetailsAttributesValueItem) agentapi.APIHproxyParseBookPostOKAttributesItemValuesItem {
					return agentapi.APIHproxyParseBookPostOKAttributesItemValuesItem{
						Name: v.Value,
						URL:  OptURL(v.URL),
					}
				}),
			}
		}),
	}, nil
}
