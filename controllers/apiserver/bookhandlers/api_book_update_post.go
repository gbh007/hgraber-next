package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookUpdatePost(ctx context.Context, req *serverapi.BookRaw) (serverapi.APIBookUpdatePostRes, error) {
	err := c.rebuilderUseCases.UpdateBook(ctx, apiservercore.ConvertBookRawToBookFull(req))

	if errors.Is(err, core.BookNotFoundError) {
		return &serverapi.APIBookUpdatePostNotFound{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookUpdatePostInternalServerError{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookUpdatePostNoContent{}, nil
}
