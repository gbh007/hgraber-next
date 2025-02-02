package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (a *Adapter) BooksCheck(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookCheckPost(ctx, &agentAPI.APIParsingBookCheckPostReq{
		Urls: urls,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.BooksCheckResult:
		return convertBooksCheckResult(typedRes), nil

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
}

func convertBooksCheckResult(checkResult *agentAPI.BooksCheckResult) []entities.AgentBookCheckResult {
	return pkg.Map(checkResult.Result, func(v agentAPI.BooksCheckResultResultItem) entities.AgentBookCheckResult {
		return entities.AgentBookCheckResult{
			URL:                v.URL,
			IsUnsupported:      v.Result == agentAPI.BooksCheckResultResultItemResultUnsupported,
			IsPossible:         v.Result == agentAPI.BooksCheckResultResultItemResultOk,
			HasError:           v.Result == agentAPI.BooksCheckResultResultItemResultError,
			PossibleDuplicates: v.PossibleDuplicates,
			ErrorReason:        v.ErrorDetails.Value,
		}
	})
}
