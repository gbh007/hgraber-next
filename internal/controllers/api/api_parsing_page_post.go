package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingPagePost(ctx context.Context, req *agentapi.APIParsingPagePostReq) (agentapi.APIParsingPagePostOK, error) {
	if c.parsingUseCases == nil {
		return agentapi.APIParsingPagePostOK{}, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	body, err := c.parsingUseCases.DownloadPage(ctx, req.BookURL, req.ImageURL)
	if err != nil {
		return agentapi.APIParsingPagePostOK{}, err
	}

	return agentapi.APIParsingPagePostOK{
		Data: body,
	}, nil
}
