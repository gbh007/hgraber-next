package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) BookParse(ctx context.Context, u url.URL) (agentmodel.AgentBookDetails, error) {
	res, err := a.rawClient.APIParsingBookPost(ctx, &agentapi.APIParsingBookPostReq{
		URL: u,
	})
	if err != nil {
		return agentmodel.AgentBookDetails{}, enrichError(err)
	}

	result := agentmodel.AgentBookDetails{
		URL:       res.URL,
		Name:      res.Name,
		PageCount: res.PageCount,
		Attributes: pkg.Map(
			res.Attributes,
			func(a agentapi.BookDetailsAttributesItem) agentmodel.AgentBookDetailsAttributesItem {
				return agentmodel.AgentBookDetailsAttributesItem{
					Code:   string(a.Code),
					Values: a.Values,
				}
			},
		),
		Pages: pkg.Map(res.Pages, func(p agentapi.BookDetailsPagesItem) agentmodel.AgentBookDetailsPagesItem {
			return agentmodel.AgentBookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
	}

	return result, nil
}
