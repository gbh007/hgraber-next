package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsGetPost(
	ctx context.Context,
	req *serverapi.APIFsGetPostReq,
) (serverapi.APIFsGetPostRes, error) {
	fs, err := c.fsUseCases.FileStorage(ctx, req.ID)
	if err != nil {
		return &serverapi.APIFsGetPostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	result := apiservercore.ConvertFileSystemInfoToAPI(fs)

	return &result, nil
}
