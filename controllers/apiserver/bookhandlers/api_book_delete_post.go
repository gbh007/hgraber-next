package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *BookHandlersController) APIBookDeletePost(ctx context.Context, req *serverAPI.APIBookDeletePostReq) (serverAPI.APIBookDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteBook(ctx, req.ID)

	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookDeletePostNotFound{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookDeletePostNoContent{}, nil
}
