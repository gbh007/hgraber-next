package adapter

import (
	"context"
	"fmt"
	"io"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (a *Adapter) PageLoad(ctx context.Context, url entities.AgentPageURL) (io.Reader, error) {
	res, err := a.rawClient.APIParsingPagePost(ctx, &agentAPI.APIParsingPagePostReq{
		BookURL:  url.BookURL,
		ImageURL: url.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIParsingPagePostOK:
		return typedRes.Data, nil

	case *agentAPI.APIParsingPagePostBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingPagePostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingPagePostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingPagePostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}
}
