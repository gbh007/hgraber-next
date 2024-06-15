package adapter

import (
	"context"
	"fmt"
	"net/url"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (a *Adapter) BooksCheck(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookCheckPost(ctx, &client.APIParsingBookCheckPostReq{
		Urls: urls,
	})
	if err != nil {
		return nil, err
	}

	var result []entities.AgentBookCheckResult

	switch typedRes := res.(type) {
	case *client.APIParsingBookCheckPostOK:
		result = pkg.Map(typedRes.Result, func(v client.APIParsingBookCheckPostOKResultItem) entities.AgentBookCheckResult {
			return entities.AgentBookCheckResult{
				URL:                v.URL,
				IsUnsupported:      v.Result == client.APIParsingBookCheckPostOKResultItemResultUnsupported,
				IsPossible:         v.Result == client.APIParsingBookCheckPostOKResultItemResultOk,
				HasError:           v.Result == client.APIParsingBookCheckPostOKResultItemResultError,
				PossibleDuplicates: v.PossibleDuplicates,
				ErrorReason:        v.ErrorDetails.Value,
			}
		})

	case *client.APIParsingBookCheckPostBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIParsingBookCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIParsingBookCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIParsingBookCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}

	return result, nil
}
