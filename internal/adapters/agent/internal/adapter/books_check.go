package adapter

import (
	"context"
	"fmt"
	"net/url"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/agentAPI"
)

func (a *Adapter) BooksCheck(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookCheckPost(ctx, &agentAPI.APIParsingBookCheckPostReq{
		Urls: urls,
	})
	if err != nil {
		return nil, err
	}

	var result []entities.AgentBookCheckResult

	switch typedRes := res.(type) {
	case *agentAPI.APIParsingBookCheckPostOK:
		result = pkg.Map(typedRes.Result, func(v agentAPI.APIParsingBookCheckPostOKResultItem) entities.AgentBookCheckResult {
			return entities.AgentBookCheckResult{
				URL:                v.URL,
				IsUnsupported:      v.Result == agentAPI.APIParsingBookCheckPostOKResultItemResultUnsupported,
				IsPossible:         v.Result == agentAPI.APIParsingBookCheckPostOKResultItemResultOk,
				HasError:           v.Result == agentAPI.APIParsingBookCheckPostOKResultItemResultError,
				PossibleDuplicates: v.PossibleDuplicates,
				ErrorReason:        v.ErrorDetails.Value,
			}
		})

	case *agentAPI.APIParsingBookCheckPostBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingBookCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingBookCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingBookCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}

	return result, nil
}
