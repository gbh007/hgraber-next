package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByBookPagesDeletePost(ctx context.Context, req *serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostReq) (serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostRes, error) {
	err := c.deduplicateUseCases.UnMarkBookPagesAsDeadHash(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostNoContent{}, nil
}
