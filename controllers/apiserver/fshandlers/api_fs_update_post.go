package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsUpdatePost(ctx context.Context, req *serverapi.FileSystemInfo) (serverapi.APIFsUpdatePostRes, error) {
	err := c.fsUseCases.UpdateFileStorage(ctx, apiservercore.ConvertFileSystemInfoFromAPI(req))
	if err != nil {
		return &serverapi.APIFsUpdatePostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsUpdatePostNoContent{}, nil
}
