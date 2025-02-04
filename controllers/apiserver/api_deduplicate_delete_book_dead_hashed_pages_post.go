package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateDeleteBookDeadHashedPagesPost(ctx context.Context, req *serverAPI.APIDeduplicateDeleteBookDeadHashedPagesPostReq) (serverAPI.APIDeduplicateDeleteBookDeadHashedPagesPostRes, error) {
	err := c.deduplicateUseCases.DeleteBookDeadHashedPages(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateDeleteBookDeadHashedPagesPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeleteBookDeadHashedPagesPostNoContent{}, nil
}
