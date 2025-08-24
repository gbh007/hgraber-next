package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeadHashSetPost(
	ctx context.Context,
	req *serverapi.APIDeduplicateDeadHashSetPostReq,
) (serverapi.APIDeduplicateDeadHashSetPostRes, error) {
	var err error

	switch {
	case req.Value && !req.PageNumber.Set:
		err = c.deduplicateUseCases.MarkBookPagesAsDeadHash(ctx, req.BookID)
	case !req.Value && !req.PageNumber.Set:
		err = c.deduplicateUseCases.UnMarkBookPagesAsDeadHash(ctx, req.BookID)
	case req.Value && req.PageNumber.Set:
		err = c.deduplicateUseCases.CreateDeadHashByPage(ctx, req.BookID, req.PageNumber.Value)
	case !req.Value && req.PageNumber.Set:
		err = c.deduplicateUseCases.DeleteDeadHashByPage(ctx, req.BookID, req.PageNumber.Value)
	}

	if err != nil {
		return &serverapi.APIDeduplicateDeadHashSetPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeadHashSetPostNoContent{}, nil
}
