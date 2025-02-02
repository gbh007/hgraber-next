package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateDeadHashByPageDeletePost(ctx context.Context, req *serverAPI.APIDeduplicateDeadHashByPageDeletePostReq) (serverAPI.APIDeduplicateDeadHashByPageDeletePostRes, error) {
	err := c.deduplicateUseCases.DeleteDeadHashByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverAPI.APIDeduplicateDeadHashByPageDeletePostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateDeadHashByPageDeletePostNoContent{}, nil
}
