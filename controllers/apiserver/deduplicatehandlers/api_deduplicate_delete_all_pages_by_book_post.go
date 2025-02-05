package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeleteAllPagesByBookPost(ctx context.Context, req *serverapi.APIDeduplicateDeleteAllPagesByBookPostReq) (serverapi.APIDeduplicateDeleteAllPagesByBookPostRes, error) {
	err := c.deduplicateUseCases.RemoveBookPagesWithDeadHash(ctx, req.BookID, req.MarkAsDeletedEmptyBook.Value)
	if err != nil {
		return &serverapi.APIDeduplicateDeleteAllPagesByBookPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeleteAllPagesByBookPostNoContent{}, nil
}
