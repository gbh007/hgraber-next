package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsTransferPost(ctx context.Context, req *serverapi.APIFsTransferPostReq) (serverapi.APIFsTransferPostRes, error) {
	err := c.fsUseCases.TransferFSFiles(ctx, req.From, req.To, req.OnlyPreviewPages.Value)
	if err != nil {
		return &serverapi.APIFsTransferPostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsTransferPostNoContent{}, nil
}
