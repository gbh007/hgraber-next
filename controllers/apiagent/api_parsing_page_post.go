package apiagent

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingPagePost(
	ctx context.Context,
	req *agentapi.APIParsingPagePostReq,
) (agentapi.APIParsingPagePostOK, error) {
	body, err := c.parsingUseCases.PageBodyByURL(ctx, req.ImageURL)
	if err != nil {
		return agentapi.APIParsingPagePostOK{}, apiError{
			Code:      http.StatusInternalServerError,
			InnerCode: ParseUseCaseCode,
			Details:   err.Error(),
		}
	}

	return agentapi.APIParsingPagePostOK{
		Data: body,
	}, nil
}
