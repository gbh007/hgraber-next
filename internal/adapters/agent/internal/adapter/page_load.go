package adapter

import (
	"context"
	"fmt"
	"io"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
)

func (a *Adapter) PageLoad(ctx context.Context, url entities.AgentPageURL) (io.Reader, error) {
	res, err := a.rawClient.APIParsingPagePost(ctx, &client.APIParsingPagePostReq{
		BookURL:  url.BookURL,
		ImageURL: url.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *client.APIParsingPagePostOK:
		return typedRes.Data, nil

	case *client.APIParsingPagePostBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIParsingPagePostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIParsingPagePostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIParsingPagePostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}
}
