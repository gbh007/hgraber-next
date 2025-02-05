package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *BookHandlersController) APIBookUpdatePost(ctx context.Context, req *serverAPI.BookRaw) (serverAPI.APIBookUpdatePostRes, error) {
	err := c.rebuilderUseCases.UpdateBook(ctx, apiservercore.ConvertBookRawToBookFull(req))

	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookUpdatePostNotFound{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookUpdatePostInternalServerError{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookUpdatePostNoContent{}, nil
}
