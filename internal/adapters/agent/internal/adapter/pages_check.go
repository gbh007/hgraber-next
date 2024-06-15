package adapter

import (
	"context"
	"fmt"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (a *Adapter) PagesCheck(ctx context.Context, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error) {
	res, err := a.rawClient.APIParsingPageCheckPost(ctx, &client.APIParsingPageCheckPostReq{
		Urls: pkg.Map(urls, func(u entities.AgentPageURL) client.APIParsingPageCheckPostReqUrlsItem {
			return client.APIParsingPageCheckPostReqUrlsItem{
				BookURL:  u.BookURL,
				ImageURL: u.ImageURL,
			}
		}),
	})
	if err != nil {
		return nil, err
	}

	var result []entities.AgentPageCheckResult

	switch typedRes := res.(type) {
	case *client.APIParsingPageCheckPostOK:
		result = pkg.Map(typedRes.Result, func(v client.APIParsingPageCheckPostOKResultItem) entities.AgentPageCheckResult {
			return entities.AgentPageCheckResult{
				BookURL:       v.BookURL,
				ImageURL:      v.ImageURL,
				IsUnsupported: v.Result == client.APIParsingPageCheckPostOKResultItemResultUnsupported,
				IsPossible:    v.Result == client.APIParsingPageCheckPostOKResultItemResultOk,
				HasError:      v.Result == client.APIParsingPageCheckPostOKResultItemResultError,
				ErrorReason:   v.ErrorDetails.Value,
			}
		})

	case *client.APIParsingPageCheckPostBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIParsingPageCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIParsingPageCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIParsingPageCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}

	return result, nil
}
