package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingBookPost(ctx context.Context, req *agentapi.APIParsingBookPostReq) (agentapi.APIParsingBookPostRes, error) {
	details, err := c.parsingUseCases.ParseBook(ctx, req.URL)
	if err != nil {
		return &agentapi.APIParsingBookPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.BookDetails{
		URL:       details.URL,
		Name:      details.Name,
		PageCount: details.PageCount,
		Attributes: pkg.Map(details.Attributes, func(attr agentmodel.AgentBookDetailsAttributesItem) agentapi.BookDetailsAttributesItem {
			return agentapi.BookDetailsAttributesItem{
				Code:   agentapi.BookDetailsAttributesItemCode(attr.Code),
				Values: attr.Values,
			}
		}),
		Pages: pkg.Map(details.Pages, func(p agentmodel.AgentBookDetailsPagesItem) agentapi.BookDetailsPagesItem {
			return agentapi.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
	}, nil
}
