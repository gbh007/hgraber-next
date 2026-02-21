package systemhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorDeletePost(
	ctx context.Context,
	req *serverapi.APIParsingMirrorDeletePostReq,
) error {
	err := c.parseUseCases.DeleteMirror(ctx, req.ID)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
