package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsDeletePost(
	ctx context.Context,
	req *serverapi.APIFsDeletePostReq,
) error {
	return c.fsUseCases.DeleteFileStorage(ctx, req.ID)
}
