package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) BooksCheck(ctx context.Context, urls []url.URL) ([]serveragentdto.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookCheckPost(ctx, &agentapi.APIParsingBookCheckPostReq{
		Urls: urls,
	})
	if err != nil {
		return nil, enrichError(err)
	}

	return convertBooksCheckResult(res), nil
}

func convertBooksCheckResult(checkResult *agentapi.BooksCheckResult) []serveragentdto.AgentBookCheckResult {
	return pkg.Map(checkResult.Result, func(v agentapi.BooksCheckResultResultItem) serveragentdto.AgentBookCheckResult {
		return serveragentdto.AgentBookCheckResult{
			URL:           v.URL,
			IsUnsupported: v.Result == agentapi.BooksCheckResultResultItemResultUnsupported,
			IsPossible:    v.Result == agentapi.BooksCheckResultResultItemResultOk,
			HasError:      v.Result == agentapi.BooksCheckResultResultItemResultError,
			ErrorReason:   v.ErrorDetails.Value,
		}
	})
}
