package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIFsCreatePost(ctx context.Context, req *serverAPI.FileSystemInfo) (serverAPI.APIFsCreatePostRes, error) {
	id, err := c.fsUseCases.NewFileStorage(ctx, convertFileSystemInfoFromAPI(req))
	if err != nil {
		return &serverAPI.APIFsCreatePostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsCreatePostOK{
		ID: id,
	}, nil
}
