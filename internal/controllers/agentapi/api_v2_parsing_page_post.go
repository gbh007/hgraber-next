package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2ParsingPagePost(
	ctx context.Context,
	req *agentapi.APIV2ParsingPagePostReq,
) (agentapi.APIV2ParsingPagePostOK, error) {
	if c.parsingUseCases == nil {
		return agentapi.APIV2ParsingPagePostOK{}, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	body, err := c.parsingUseCases.Image(ctx, req.BookURL, req.ImageURL)
	if err != nil {
		return agentapi.APIV2ParsingPagePostOK{}, err
	}

	return agentapi.APIV2ParsingPagePostOK{
		Data: body,
	}, nil
}
