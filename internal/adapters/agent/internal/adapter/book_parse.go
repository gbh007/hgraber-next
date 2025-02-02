package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (a *Adapter) BookParse(ctx context.Context, url url.URL) (entities.AgentBookDetails, error) {
	res, err := a.rawClient.APIParsingBookPost(ctx, &agentAPI.APIParsingBookPostReq{
		URL: url,
	})
	if err != nil {
		return entities.AgentBookDetails{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.BookDetails:
		result := entities.AgentBookDetails{
			URL:       typedRes.URL,
			Name:      typedRes.Name,
			PageCount: typedRes.PageCount,
			Attributes: pkg.Map(typedRes.Attributes, func(a agentAPI.BookDetailsAttributesItem) entities.AgentBookDetailsAttributesItem {
				return entities.AgentBookDetailsAttributesItem{
					Code:   string(a.Code),
					Values: a.Values,
				}
			}),
			Pages: pkg.Map(typedRes.Pages, func(p agentAPI.BookDetailsPagesItem) entities.AgentBookDetailsPagesItem {
				return entities.AgentBookDetailsPagesItem{
					PageNumber: p.PageNumber,
					URL:        p.URL,
					Filename:   p.Filename,
				}
			}),
		}

		return result, nil

	case *agentAPI.APIParsingBookPostBadRequest:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingBookPostUnauthorized:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingBookPostForbidden:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingBookPostInternalServerError:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentBookDetails{}, entities.AgentAPIUnknownResponse
	}
}
