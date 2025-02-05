package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByBookPagesDeletePost(ctx context.Context, req *serverapi.APIDeduplicateDeadHashByBookPagesDeletePostReq) (serverapi.APIDeduplicateDeadHashByBookPagesDeletePostRes, error) {
	err := c.deduplicateUseCases.UnMarkBookPagesAsDeadHash(ctx, req.BookID)
	if err != nil {
		return &serverapi.APIDeduplicateDeadHashByBookPagesDeletePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeadHashByBookPagesDeletePostNoContent{}, nil
}
