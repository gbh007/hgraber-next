package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsRemoveMismatchPost(
	ctx context.Context,
	req *serverapi.APIFsRemoveMismatchPostReq,
) error {
	return c.systemUseCases.RemoveFilesInFSMismatch(ctx, req.ID)
}
