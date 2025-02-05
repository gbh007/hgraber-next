package adapter

import (
	"context"
	"fmt"
	"io"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) PageLoad(ctx context.Context, url agentmodel.AgentPageURL) (io.Reader, error) {
	res, err := a.rawClient.APIParsingPagePost(ctx, &agentapi.APIParsingPagePostReq{
		BookURL:  url.BookURL,
		ImageURL: url.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIParsingPagePostOK:
		return typedRes.Data, nil

	case *agentapi.APIParsingPagePostBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIParsingPagePostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIParsingPagePostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIParsingPagePostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}
