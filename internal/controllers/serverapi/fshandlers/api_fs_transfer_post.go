package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsTransferPost(
	ctx context.Context,
	req *serverapi.APIFsTransferPostReq,
) error {
	return c.fsUseCases.TransferFSFiles(ctx, req.From, req.To, req.OnlyPreviewPages.Value)
}
