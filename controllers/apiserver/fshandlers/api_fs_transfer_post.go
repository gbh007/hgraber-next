package fshandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsTransferPost(
	ctx context.Context,
	req *serverapi.APIFsTransferPostReq,
) error {
	err := c.fsUseCases.TransferFSFiles(ctx, req.From, req.To, req.OnlyPreviewPages.Value)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
