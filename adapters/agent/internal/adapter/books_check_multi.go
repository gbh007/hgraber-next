package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (a *Adapter) BooksCheckMulti(ctx context.Context, u url.URL) ([]agentmodel.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookMultiPost(ctx, &agentAPI.APIParsingBookMultiPostReq{
		URL: u,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.BooksCheckResult:
		return convertBooksCheckResult(typedRes), nil

	case *agentAPI.APIParsingBookMultiPostBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingBookMultiPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingBookMultiPostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingBookMultiPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}
