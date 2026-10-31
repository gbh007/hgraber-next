package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) BookParse(ctx context.Context, u url.URL) (serveragentdto.AgentBookDetails, error) {
	res, err := a.rawClient.APIParsingBookPost(ctx, &agentapi.APIParsingBookPostReq{
		URL: u,
	})
	if err != nil {
		return serveragentdto.AgentBookDetails{}, enrichError(err)
	}

	result := serveragentdto.AgentBookDetails{
		URL:       res.URL,
		Name:      res.Name,
		PageCount: res.PageCount,
		Attributes: pkg.Map(
			res.Attributes,
			func(a agentapi.BookDetailsAttributesItem) serveragentdto.AgentBookDetailsAttributesItem {
				return serveragentdto.AgentBookDetailsAttributesItem{
					Code:   string(a.Code),
					Values: a.Values,
				}
			},
		),
		Pages: pkg.Map(res.Pages, func(p agentapi.BookDetailsPagesItem) serveragentdto.AgentBookDetailsPagesItem {
			return serveragentdto.AgentBookDetailsPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
	}

	return result, nil
}
