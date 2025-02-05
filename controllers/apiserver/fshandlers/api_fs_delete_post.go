package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsDeletePost(ctx context.Context, req *serverapi.APIFsDeletePostReq) (serverapi.APIFsDeletePostRes, error) {
	err := c.fsUseCases.DeleteFileStorage(ctx, req.ID)
	if err != nil {
		return &serverapi.APIFsDeletePostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsDeletePostNoContent{}, nil
}
