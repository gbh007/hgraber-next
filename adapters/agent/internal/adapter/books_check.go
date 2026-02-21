package adapter

import (
	"context"
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
		return nil, enrichError(err)
	}

	return convertBooksCheckResult(res), nil
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
