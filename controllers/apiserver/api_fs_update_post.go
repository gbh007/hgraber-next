package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIFsUpdatePost(ctx context.Context, req *serverAPI.FileSystemInfo) (serverAPI.APIFsUpdatePostRes, error) {
	err := c.fsUseCases.UpdateFileStorage(ctx, convertFileSystemInfoFromAPI(req))
	if err != nil {
		return &serverAPI.APIFsUpdatePostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsUpdatePostNoContent{}, nil
}
