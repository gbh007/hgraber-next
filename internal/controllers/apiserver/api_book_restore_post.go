package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIBookRestorePost(ctx context.Context, req *serverAPI.APIBookRestorePostReq) (serverAPI.APIBookRestorePostRes, error) {
	err := c.rebuilderUseCases.RestoreBook(ctx, req.BookID, req.OnlyPages.Value)
	if err != nil {
		return &serverAPI.APIBookRestorePostInternalServerError{
			InnerCode: RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookRestorePostNoContent{}, nil
}
