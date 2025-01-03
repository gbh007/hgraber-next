package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookRebuildPost(ctx context.Context, req *serverAPI.APIBookRebuildPostReq) (serverAPI.APIBookRebuildPostRes, error) {
	id, err := c.rebuilderUseCases.RebuildBook(ctx, entities.RebuildBookRequest{
		OldBook:         convertBookRawToBookFull(&req.OldBook),
		SelectedPages:   req.SelectedPages,
		MergeWithBook:   req.MergeWithBook.Value,
		OnlyUniquePages: req.OnlyUnique.Value,
	})

	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookRebuildPostNotFound{
			InnerCode: RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookRebuildPostInternalServerError{
			InnerCode: RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookRebuildPostOK{
		ID: id,
	}, nil
}
