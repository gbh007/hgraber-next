package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) BooksCheckMulti(ctx context.Context, u url.URL) ([]agentmodel.AgentBookCheckResult, error) {
	res, err := a.rawClient.APIParsingBookMultiPost(ctx, &agentapi.APIParsingBookMultiPostReq{
		URL: u,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentapi.BooksCheckResult:
		return convertBooksCheckResult(typedRes), nil

	case *agentapi.APIParsingBookMultiPostBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIParsingBookMultiPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIParsingBookMultiPostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIParsingBookMultiPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}
