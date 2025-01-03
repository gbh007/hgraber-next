package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateDeleteAllPagesByHashPost(ctx context.Context, req *serverAPI.APIDeduplicateDeleteAllPagesByHashPostReq) (serverAPI.APIDeduplicateDeleteAllPagesByHashPostRes, error) {
	err := c.deduplicateUseCases.DeleteAllPageByHash(ctx, req.BookID, req.PageNumber, req.SetDeadHash.Value)
	if err != nil {
		return &serverAPI.APIDeduplicateDeleteAllPagesByHashPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeleteAllPagesByHashPostNoContent{}, nil
}
