package fshandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsValidatePost(
	ctx context.Context,
	req *serverapi.APIFsValidatePostReq,
) error {
	err := c.fsUseCases.ValidateFS(ctx, req.ID)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
