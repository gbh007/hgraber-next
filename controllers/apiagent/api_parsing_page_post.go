package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingPagePost(
	ctx context.Context,
	req *agentapi.APIParsingPagePostReq,
) (agentapi.APIParsingPagePostOK, error) {
	body, err := c.parsingUseCases.PageBodyByURL(ctx, req.ImageURL)
	if err != nil {
		return agentapi.APIParsingPagePostOK{}, err //nolint:wrapcheck // будет исправлено позднее
	}

	return agentapi.APIParsingPagePostOK{
		Data: body,
	}, nil
}
