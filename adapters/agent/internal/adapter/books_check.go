package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) BooksCheck(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error) {
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
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingBookCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingBookCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingBookCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}

func convertBooksCheckResult(checkResult *agentAPI.BooksCheckResult) []agentmodel.AgentBookCheckResult {
	return pkg.Map(checkResult.Result, func(v agentAPI.BooksCheckResultResultItem) agentmodel.AgentBookCheckResult {
		return agentmodel.AgentBookCheckResult{
			URL:                v.URL,
			IsUnsupported:      v.Result == agentAPI.BooksCheckResultResultItemResultUnsupported,
			IsPossible:         v.Result == agentAPI.BooksCheckResultResultItemResultOk,
			HasError:           v.Result == agentAPI.BooksCheckResultResultItemResultError,
			PossibleDuplicates: v.PossibleDuplicates,
			ErrorReason:        v.ErrorDetails.Value,
		}
	})
}
