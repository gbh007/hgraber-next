package apiagent

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/agentAPI"
)

func (c *Controller) APIParsingBookPost(ctx context.Context, req *agentAPI.APIParsingBookPostReq) (agentAPI.APIParsingBookPostRes, error) {
	details, err := c.parsingUseCases.ParseBook(ctx, req.URL)
	if err != nil {
		return &agentAPI.APIParsingBookPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	return &agentAPI.BookDetails{
		URL:       details.URL,
		Name:      details.Name,
		PageCount: details.PageCount,
		Attributes: pkg.Map(details.Attributes, func(attr entities.AgentBookDetailsAttributesItem) agentAPI.BookDetailsAttributesItem {
			return agentAPI.BookDetailsAttributesItem{
				Code:   agentAPI.BookDetailsAttributesItemCode(attr.Code),
				Values: attr.Values,
			}
		}),
		Pages: pkg.Map(details.Pages, func(p entities.AgentBookDetailsPagesItem) agentAPI.BookDetailsPagesItem {
			return agentAPI.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
	}, nil
}
