package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateDeadHashByBookPagesDeletePost(ctx context.Context, req *serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostReq) (serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostRes, error) {
	err := c.deduplicateUseCases.UnMarkBookPagesAsDeadHash(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeadHashByBookPagesDeletePostNoContent{}, nil
}
