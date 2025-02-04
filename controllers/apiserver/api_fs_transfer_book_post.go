package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIFsTransferBookPost(ctx context.Context, req *serverAPI.APIFsTransferBookPostReq) (serverAPI.APIFsTransferBookPostRes, error) {
	var pageNumber *int

	if req.OnlyPreviewPages.Value {
		p := core.PageNumberForPreview
		pageNumber = &p
	}

	if req.PageNumber.IsSet() {
		pageNumber = &req.PageNumber.Value
	}

	err := c.fsUseCases.TransferFSFilesByBook(ctx, req.BookID, req.To, pageNumber)
	if err != nil {
		return &serverAPI.APIFsTransferBookPostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsTransferBookPostNoContent{}, nil
}
