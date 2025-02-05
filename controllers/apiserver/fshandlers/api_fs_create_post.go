package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsCreatePost(ctx context.Context, req *serverapi.FileSystemInfo) (serverapi.APIFsCreatePostRes, error) {
	id, err := c.fsUseCases.NewFileStorage(ctx, apiservercore.ConvertFileSystemInfoFromAPI(req))
	if err != nil {
		return &serverapi.APIFsCreatePostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsCreatePostOK{
		ID: id,
	}, nil
}
