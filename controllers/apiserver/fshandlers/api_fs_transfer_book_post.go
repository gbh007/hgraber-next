package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsTransferBookPost(
	ctx context.Context,
	req *serverapi.APIFsTransferBookPostReq,
) error {
	var pageNumber *int

	if req.OnlyPreviewPages.Value {
		p := core.PageNumberForPreview
		pageNumber = &p
	}

	if req.PageNumber.IsSet() {
		pageNumber = &req.PageNumber.Value
	}

	return c.fsUseCases.TransferFSFilesByBook(ctx, req.BookID, req.To, pageNumber)
}
