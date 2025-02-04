package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIFsDeletePost(ctx context.Context, req *serverAPI.APIFsDeletePostReq) (serverAPI.APIFsDeletePostRes, error) {
	err := c.fsUseCases.DeleteFileStorage(ctx, req.ID)
	if err != nil {
		return &serverAPI.APIFsDeletePostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsDeletePostNoContent{}, nil
}
