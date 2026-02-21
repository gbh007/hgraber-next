package bookhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookRestorePost(
	ctx context.Context,
	req *serverapi.APIBookRestorePostReq,
) error {
	err := c.rebuilderUseCases.RestoreBook(ctx, req.BookID, req.OnlyPages.Value)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
