package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingPagePost(ctx context.Context, req *agentapi.APIParsingPagePostReq) (agentapi.APIParsingPagePostRes, error) {
	body, err := c.parsingUseCases.DownloadPage(ctx, req.BookURL, req.ImageURL)
	if err != nil {
		return &agentapi.APIParsingPagePostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIParsingPagePostOK{
		Data: body,
	}, nil
}
