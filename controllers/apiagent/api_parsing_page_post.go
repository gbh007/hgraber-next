package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (c *Controller) APIParsingPagePost(ctx context.Context, req *agentAPI.APIParsingPagePostReq) (agentAPI.APIParsingPagePostRes, error) {
	body, err := c.parsingUseCases.DownloadPage(ctx, req.BookURL, req.ImageURL)
	if err != nil {
		return &agentAPI.APIParsingPagePostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	return &agentAPI.APIParsingPagePostOK{
		Data: body,
	}, nil
}
