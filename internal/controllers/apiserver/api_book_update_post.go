package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIBookUpdatePost(ctx context.Context, req *serverAPI.BookRaw) (serverAPI.APIBookUpdatePostRes, error) {
	err := c.rebuilderUseCases.UpdateBook(ctx, convertBookRawToBookFull(req))

	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookUpdatePostNotFound{
			InnerCode: RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookUpdatePostInternalServerError{
			InnerCode: RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookUpdatePostNoContent{}, nil
}
