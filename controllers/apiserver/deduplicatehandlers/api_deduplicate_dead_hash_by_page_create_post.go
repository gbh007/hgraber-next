package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByPageCreatePost(ctx context.Context, req *serverapi.APIDeduplicateDeadHashByPageCreatePostReq) (serverapi.APIDeduplicateDeadHashByPageCreatePostRes, error) {
	err := c.deduplicateUseCases.CreateDeadHashByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverapi.APIDeduplicateDeadHashByPageCreatePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeadHashByPageCreatePostNoContent{}, nil
}
