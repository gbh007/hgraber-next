package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByBookPagesCreatePost(ctx context.Context, req *serverAPI.APIDeduplicateDeadHashByBookPagesCreatePostReq) (serverAPI.APIDeduplicateDeadHashByBookPagesCreatePostRes, error) {
	err := c.deduplicateUseCases.MarkBookPagesAsDeadHash(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateDeadHashByBookPagesCreatePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeadHashByBookPagesCreatePostNoContent{}, nil
}
