package adapter

import (
	"context"
	"fmt"
	"net/url"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (a *Adapter) BookParse(ctx context.Context, url url.URL) (entities.AgentBookDetails, error) {
	res, err := a.rawClient.APIParsingBookPost(ctx, &client.APIParsingBookPostReq{
		URL: url,
	})
	if err != nil {
		return entities.AgentBookDetails{}, err
	}

	switch typedRes := res.(type) {
	case *client.BookDetails:
		result := entities.AgentBookDetails{
			URL:       typedRes.URL,
			Name:      typedRes.Name,
			PageCount: typedRes.PageCount,
			Attributes: pkg.Map(typedRes.Attributes, func(a client.BookDetailsAttributesItem) entities.AgentBookDetailsAttributesItem {
				return entities.AgentBookDetailsAttributesItem{
					Code:   string(a.Code),
					Values: a.Values,
				}
			}),
			Pages: pkg.Map(typedRes.Pages, func(p client.BookDetailsPagesItem) entities.AgentBookDetailsPagesItem {
				return entities.AgentBookDetailsPagesItem{
					PageNumber: p.PageNumber,
					URL:        p.URL,
					Filename:   p.Filename,
				}
			}),
		}

		return result, nil

	case *client.APIParsingBookPostBadRequest:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIParsingBookPostUnauthorized:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIParsingBookPostForbidden:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIParsingBookPostInternalServerError:
		return entities.AgentBookDetails{}, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentBookDetails{}, entities.AgentAPIUnknownResponse
	}
}
