package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeleteAllPagesByHashPost(ctx context.Context, req *serverapi.APIDeduplicateDeleteAllPagesByHashPostReq) (serverapi.APIDeduplicateDeleteAllPagesByHashPostRes, error) {
	err := c.deduplicateUseCases.DeleteAllPageByHash(ctx, req.BookID, req.PageNumber, req.SetDeadHash.Value)
	if err != nil {
		return &serverapi.APIDeduplicateDeleteAllPagesByHashPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeleteAllPagesByHashPostNoContent{}, nil
}
