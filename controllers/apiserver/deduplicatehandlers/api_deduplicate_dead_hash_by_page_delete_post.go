package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByPageDeletePost(ctx context.Context, req *serverAPI.APIDeduplicateDeadHashByPageDeletePostReq) (serverAPI.APIDeduplicateDeadHashByPageDeletePostRes, error) {
	err := c.deduplicateUseCases.DeleteDeadHashByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverAPI.APIDeduplicateDeadHashByPageDeletePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeadHashByPageDeletePostNoContent{}, nil
}
