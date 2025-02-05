package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIFsTransferPost(ctx context.Context, req *serverAPI.APIFsTransferPostReq) (serverAPI.APIFsTransferPostRes, error) {
	err := c.fsUseCases.TransferFSFiles(ctx, req.From, req.To, req.OnlyPreviewPages.Value)
	if err != nil {
		return &serverAPI.APIFsTransferPostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsTransferPostNoContent{}, nil
}
