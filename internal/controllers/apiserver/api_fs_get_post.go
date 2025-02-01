package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIFsGetPost(ctx context.Context, req *serverAPI.APIFsGetPostReq) (serverAPI.APIFsGetPostRes, error) {
	fs, err := c.fsUseCases.FileStorage(ctx, req.ID)
	if err != nil {
		return &serverAPI.APIFsGetPostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	result := convertFileSystemInfoToAPI(fs)

	return &result, nil
}
