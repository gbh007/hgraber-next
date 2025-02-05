package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *BookHandlersController) APIBookRebuildPost(ctx context.Context, req *serverAPI.APIBookRebuildPostReq) (serverAPI.APIBookRebuildPostRes, error) {
	id, err := c.rebuilderUseCases.RebuildBook(ctx, core.RebuildBookRequest{
		ModifiedOldBook: apiservercore.ConvertBookRawToBookFull(&req.OldBook),
		SelectedPages:   req.SelectedPages,
		MergeWithBook:   req.MergeWithBook.Value,
		PageOrder:       req.PageOrder,
		Flags: core.RebuildBookRequestFlags{
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

	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookRebuildPostNotFound{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookRebuildPostInternalServerError{
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookRebuildPostOK{
		ID: id,
	}, nil
}
