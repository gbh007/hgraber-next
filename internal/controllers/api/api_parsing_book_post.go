package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingBookPost(ctx context.Context, req *agentapi.APIParsingBookPostReq) (*agentapi.BookDetails, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	details, err := c.parsingUseCases.ParseBook(ctx, req.URL)
	if err != nil {
		return nil, err
	}

	return &agentapi.BookDetails{
		URL:       details.URL,
		Name:      details.Name,
		PageCount: details.PageCount,
		Attributes: pkg.Map(details.Attributes, func(attr entities.AgentBookDetailsAttributesItem) agentapi.BookDetailsAttributesItem {
			return agentapi.BookDetailsAttributesItem{
				Code:   agentapi.BookDetailsAttributesItemCode(attr.Code),
				Values: attr.Values,
			}
		}),
		Pages: pkg.Map(details.Pages, func(p entities.AgentBookDetailsPagesItem) agentapi.BookDetailsPagesItem {
			return agentapi.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
	}, nil
}
