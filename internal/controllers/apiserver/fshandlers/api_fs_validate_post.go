package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsValidatePost(
	ctx context.Context,
	req *serverapi.APIFsValidatePostReq,
) error {
	return c.fsUseCases.ValidateFS(ctx, req.ID)
}
