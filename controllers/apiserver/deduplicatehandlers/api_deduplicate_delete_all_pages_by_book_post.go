package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeleteAllPagesByBookPost(ctx context.Context, req *serverAPI.APIDeduplicateDeleteAllPagesByBookPostReq) (serverAPI.APIDeduplicateDeleteAllPagesByBookPostRes, error) {
	err := c.deduplicateUseCases.RemoveBookPagesWithDeadHash(ctx, req.BookID, req.MarkAsDeletedEmptyBook.Value)
	if err != nil {
		return &serverAPI.APIDeduplicateDeleteAllPagesByBookPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeleteAllPagesByBookPostNoContent{}, nil
}
