package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsRemoveMismatchPost(ctx context.Context, req *serverapi.APIFsRemoveMismatchPostReq) (serverapi.APIFsRemoveMismatchPostRes, error) {
	err := c.taskUseCases.RemoveFilesInFSMismatch(ctx, req.ID)
	if err != nil {
		return &serverapi.APIFsRemoveMismatchPostInternalServerError{
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsRemoveMismatchPostNoContent{}, nil
}
