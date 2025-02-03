package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIBookRebuildPost(ctx context.Context, req *serverAPI.APIBookRebuildPostReq) (serverAPI.APIBookRebuildPostRes, error) {
	id, err := c.rebuilderUseCases.RebuildBook(ctx, entities.RebuildBookRequest{
		ModifiedOldBook: convertBookRawToBookFull(&req.OldBook),
		SelectedPages:   req.SelectedPages,
		MergeWithBook:   req.MergeWithBook.Value,
		PageOrder:       req.PageOrder,
		Flags: entities.RebuildBookRequestFlags{
			OnlyUniquePages:      req.Flags.Value.OnlyUnique.Value,
			ExcludeDeadHashPages: req.Flags.Value.ExcludeDeadHashPages.Value,
			Only1CopyPages:       req.Flags.Value.Only1Copy.Value,

			SetOriginLabels: req.Flags.Value.SetOriginLabels.Value,
			AutoVerify:      req.Flags.Value.AutoVerify.Value,

			ExtractMode: req.Flags.Value.ExtractMode.Value,

			PageReOrder: req.Flags.Value.PageReOrder.Value,

			MarkUnusedPagesAsDeadHash:              req.Flags.Value.MarkUnusedPagesAsDeadHash.Value,
			MarkUnusedPagesAsDeleted:               req.Flags.Value.MarkUnusedPagesAsDeleted.Value,
			MarkEmptyBookAsDeletedAfterRemovePages: req.Flags.Value.MarkEmptyBookAsDeletedAfterRemovePages.Value,
		},
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
