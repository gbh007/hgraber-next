package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookDeletePost(ctx context.Context, req *serverapi.APIBookDeletePostReq) (serverapi.APIBookDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteBook(ctx, req.ID)

	if errors.Is(err, core.BookNotFoundError) {
		return &serverapi.APIBookDeletePostNotFound{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookDeletePostNoContent{}, nil
}
