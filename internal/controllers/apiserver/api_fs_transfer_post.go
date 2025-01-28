package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIFsTransferPost(ctx context.Context, req *serverAPI.APIFsTransferPostReq) (serverAPI.APIFsTransferPostRes, error) {
	err := c.fsUseCases.TransferFSFiles(ctx, req.From, req.To, req.OnlyPreviewPages.Value)
	if err != nil {
		return &serverAPI.APIFsTransferPostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsTransferPostNoContent{}, nil
}
