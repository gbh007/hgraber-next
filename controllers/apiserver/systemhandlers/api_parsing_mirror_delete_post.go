package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorDeletePost(ctx context.Context, req *serverapi.APIParsingMirrorDeletePostReq) (serverapi.APIParsingMirrorDeletePostRes, error) {
	err := c.parseUseCases.DeleteMirror(ctx, req.ID)
	if err != nil {
		return &serverapi.APIParsingMirrorDeletePostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingMirrorDeletePostNoContent{}, nil
}
