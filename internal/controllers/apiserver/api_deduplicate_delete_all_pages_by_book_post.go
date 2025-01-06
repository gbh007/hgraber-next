package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateDeleteAllPagesByBookPost(ctx context.Context, req *serverAPI.APIDeduplicateDeleteAllPagesByBookPostReq) (serverAPI.APIDeduplicateDeleteAllPagesByBookPostRes, error) {
	err := c.deduplicateUseCases.RemoveBookPagesWithDeadHash(ctx, req.BookID, req.MarkAsDeletedEmptyBook.Value)
	if err != nil {
		return &serverAPI.APIDeduplicateDeleteAllPagesByBookPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeleteAllPagesByBookPostNoContent{}, nil
}
