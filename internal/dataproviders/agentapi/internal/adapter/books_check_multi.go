package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) BooksCheckMulti(ctx context.Context, u url.URL) ([]serveragentdto.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookMultiPost(ctx, &agentapi.APIParsingBookMultiPostReq{
		URL: u,
	})
	if err != nil {
		return nil, enrichError(err)
	}

	return convertBooksCheckResult(res), nil
}
