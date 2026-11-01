package agentcacheapi

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2ParsingPagePost(
	ctx context.Context,
	req *agentapi.APIV2ParsingPagePostReq,
) (agentapi.APIV2ParsingPagePostOK, error) {
	body, err := c.parsingUseCases.PageBodyByURL(ctx, req.ImageURL)
	if err != nil {
		return agentapi.APIV2ParsingPagePostOK{}, err //nolint:wrapcheck // будет исправлено позднее
	}

	return agentapi.APIV2ParsingPagePostOK{
		Data: body,
	}, nil
}
