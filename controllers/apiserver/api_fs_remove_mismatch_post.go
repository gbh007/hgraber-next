package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIFsRemoveMismatchPost(ctx context.Context, req *serverAPI.APIFsRemoveMismatchPostReq) (serverAPI.APIFsRemoveMismatchPostRes, error) {
	err := c.taskUseCases.RemoveFilesInFSMismatch(ctx, req.ID)
	if err != nil {
		return &serverAPI.APIFsRemoveMismatchPostInternalServerError{
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsRemoveMismatchPostNoContent{}, nil
}
