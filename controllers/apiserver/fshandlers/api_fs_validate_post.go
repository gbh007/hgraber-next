package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsValidatePost(ctx context.Context, req *serverapi.APIFsValidatePostReq) (serverapi.APIFsValidatePostRes, error) {
	err := c.fsUseCases.ValidateFS(ctx, req.ID)
	if err != nil {
		return &serverapi.APIFsValidatePostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsValidatePostNoContent{}, nil
}
