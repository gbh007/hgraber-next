package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByBookPagesCreatePost(ctx context.Context, req *serverapi.APIDeduplicateDeadHashByBookPagesCreatePostReq) (serverapi.APIDeduplicateDeadHashByBookPagesCreatePostRes, error) {
	err := c.deduplicateUseCases.MarkBookPagesAsDeadHash(ctx, req.BookID)
	if err != nil {
		return &serverapi.APIDeduplicateDeadHashByBookPagesCreatePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeadHashByBookPagesCreatePostNoContent{}, nil
}
