package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) BookParse(ctx context.Context, url url.URL) (agentmodel.AgentBookDetails, error) {
	res, err := a.rawClient.APIParsingBookPost(ctx, &agentAPI.APIParsingBookPostReq{
		URL: url,
	})
	if err != nil {
		return agentmodel.AgentBookDetails{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.BookDetails:
		result := agentmodel.AgentBookDetails{
			URL:       typedRes.URL,
			Name:      typedRes.Name,
			PageCount: typedRes.PageCount,
			Attributes: pkg.Map(typedRes.Attributes, func(a agentAPI.BookDetailsAttributesItem) agentmodel.AgentBookDetailsAttributesItem {
				return agentmodel.AgentBookDetailsAttributesItem{
					Code:   string(a.Code),
					Values: a.Values,
				}
			}),
			Pages: pkg.Map(typedRes.Pages, func(p agentAPI.BookDetailsPagesItem) agentmodel.AgentBookDetailsPagesItem {
				return agentmodel.AgentBookDetailsPagesItem{
					PageNumber: p.PageNumber,
					URL:        p.URL,
					Filename:   p.Filename,
				}
			}),
		}

		return result, nil

	case *agentAPI.APIParsingBookPostBadRequest:
		return agentmodel.AgentBookDetails{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingBookPostUnauthorized:
		return agentmodel.AgentBookDetails{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingBookPostForbidden:
		return agentmodel.AgentBookDetails{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingBookPostInternalServerError:
		return agentmodel.AgentBookDetails{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.AgentBookDetails{}, agentmodel.AgentAPIUnknownResponse
	}
}
