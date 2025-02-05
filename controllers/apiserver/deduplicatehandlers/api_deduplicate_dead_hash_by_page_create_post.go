package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByPageCreatePost(ctx context.Context, req *serverAPI.APIDeduplicateDeadHashByPageCreatePostReq) (serverAPI.APIDeduplicateDeadHashByPageCreatePostRes, error) {
	err := c.deduplicateUseCases.CreateDeadHashByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverAPI.APIDeduplicateDeadHashByPageCreatePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeadHashByPageCreatePostNoContent{}, nil
}
