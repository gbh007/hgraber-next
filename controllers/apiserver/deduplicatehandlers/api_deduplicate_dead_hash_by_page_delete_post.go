package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashByPageDeletePost(ctx context.Context, req *serverapi.APIDeduplicateDeadHashByPageDeletePostReq) (serverapi.APIDeduplicateDeadHashByPageDeletePostRes, error) {
	err := c.deduplicateUseCases.DeleteDeadHashByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverapi.APIDeduplicateDeadHashByPageDeletePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeadHashByPageDeletePostNoContent{}, nil
}
