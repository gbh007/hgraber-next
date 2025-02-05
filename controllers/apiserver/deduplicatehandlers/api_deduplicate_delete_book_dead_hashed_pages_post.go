package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateDeleteBookDeadHashedPagesPost(ctx context.Context, req *serverapi.APIDeduplicateDeleteBookDeadHashedPagesPostReq) (serverapi.APIDeduplicateDeleteBookDeadHashedPagesPostRes, error) {
	err := c.deduplicateUseCases.DeleteBookDeadHashedPages(ctx, req.BookID)
	if err != nil {
		return &serverapi.APIDeduplicateDeleteBookDeadHashedPagesPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateDeleteBookDeadHashedPagesPostNoContent{}, nil
}
