package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *FSHandlersController) APIFsUpdatePost(ctx context.Context, req *serverAPI.FileSystemInfo) (serverAPI.APIFsUpdatePostRes, error) {
	err := c.fsUseCases.UpdateFileStorage(ctx, apiservercore.ConvertFileSystemInfoFromAPI(req))
	if err != nil {
		return &serverAPI.APIFsUpdatePostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsUpdatePostNoContent{}, nil
}
