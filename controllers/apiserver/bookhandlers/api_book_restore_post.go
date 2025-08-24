package bookhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookRestorePost(
	ctx context.Context,
	req *serverapi.APIBookRestorePostReq,
) (serverapi.APIBookRestorePostRes, error) {
	err := c.rebuilderUseCases.RestoreBook(ctx, req.BookID, req.OnlyPages.Value)
	if err != nil {
		return &serverapi.APIBookRestorePostInternalServerError{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookRestorePostNoContent{}, nil
}
