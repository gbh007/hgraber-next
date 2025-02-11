package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) BooksCheck(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookCheckPost(ctx, &agentapi.APIParsingBookCheckPostReq{
		Urls: urls,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentapi.BooksCheckResult:
		return convertBooksCheckResult(typedRes), nil

	case *agentapi.APIParsingBookCheckPostBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIParsingBookCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIParsingBookCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIParsingBookCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}

func convertBooksCheckResult(checkResult *agentapi.BooksCheckResult) []agentmodel.AgentBookCheckResult {
	return pkg.Map(checkResult.Result, func(v agentapi.BooksCheckResultResultItem) agentmodel.AgentBookCheckResult {
		return agentmodel.AgentBookCheckResult{
			URL:           v.URL,
			IsUnsupported: v.Result == agentapi.BooksCheckResultResultItemResultUnsupported,
			IsPossible:    v.Result == agentapi.BooksCheckResultResultItemResultOk,
			HasError:      v.Result == agentapi.BooksCheckResultResultItemResultError,
			ErrorReason:   v.ErrorDetails.Value,
		}
	})
}
