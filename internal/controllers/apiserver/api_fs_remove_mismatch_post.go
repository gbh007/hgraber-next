package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIFsRemoveMismatchPost(ctx context.Context, req *serverAPI.APIFsRemoveMismatchPostReq) (serverAPI.APIFsRemoveMismatchPostRes, error) {
	err := c.taskUseCases.RemoveFilesInFSMismatch(ctx, req.ID)
	if err != nil {
		return &serverAPI.APIFsRemoveMismatchPostInternalServerError{
			InnerCode: TaskerUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsRemoveMismatchPostNoContent{}, nil
}
