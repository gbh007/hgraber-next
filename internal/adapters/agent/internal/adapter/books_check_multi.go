package adapter

import (
	"context"
	"fmt"
	"net/url"

	"hgnext/internal/entities"
	"hgnext/open_api/agentAPI"
)

func (a *Adapter) BooksCheckMulti(ctx context.Context, u url.URL) ([]entities.AgentBookCheckResult, error) {
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
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingBookMultiPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingBookMultiPostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingBookMultiPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}
}
