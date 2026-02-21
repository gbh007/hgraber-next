package fshandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsRemoveMismatchPost(
	ctx context.Context,
	req *serverapi.APIFsRemoveMismatchPostReq,
) error {
	err := c.systemUseCases.RemoveFilesInFSMismatch(ctx, req.ID)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
