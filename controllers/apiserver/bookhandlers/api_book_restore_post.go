package bookhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookRestorePost(
	ctx context.Context,
	req *serverapi.APIBookRestorePostReq,
) error {
	return c.rebuilderUseCases.RestoreBook(ctx, req.BookID, req.OnlyPages.Value)
}
