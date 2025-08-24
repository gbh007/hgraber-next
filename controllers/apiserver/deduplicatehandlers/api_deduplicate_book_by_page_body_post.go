package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateBookByPageBodyPost(
	ctx context.Context,
	req *serverapi.APIDeduplicateBookByPageBodyPostReq,
) (serverapi.APIDeduplicateBookByPageBodyPostRes, error) {
	data, err := c.deduplicateUseCases.BookByPageEntryPercentage(ctx, req.BookID)
	if err != nil {
		return &serverapi.APIDeduplicateBookByPageBodyPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateBookByPageBodyPostOK{
		Result: pkg.Map(
			data,
			func(raw bff.DeduplicateBookResult) serverapi.APIDeduplicateBookByPageBodyPostOKResultItem {
				return serverapi.APIDeduplicateBookByPageBodyPostOKResultItem{
					Book:                                 c.apiCore.ConvertSimpleBook(raw.TargetBook, raw.PreviewPage),
					OriginCoveredTarget:                  raw.EntryPercentage,
					TargetCoveredOrigin:                  raw.ReverseEntryPercentage,
					OriginCoveredTargetWithoutDeadHashes: raw.EntryPercentageWithoutDeadHashes,
					TargetCoveredOriginWithoutDeadHashes: raw.ReverseEntryPercentageWithoutDeadHashes,

					TargetSize:                       raw.TargetSize.Size,
					SharedSize:                       raw.SharedSize,
					SharedSizeWithoutDeadHashes:      raw.SharedSizeWithoutDeadHashes,
					SharedPageCount:                  raw.SharedPages,
					SharedPageCountWithoutDeadHashes: raw.SharedPagesWithoutDeadHashes,

					TargetSizeFormatted:                  core.PrettySize(raw.TargetSize.Size),
					SharedSizeFormatted:                  core.PrettySize(raw.SharedSize),
					SharedSizeWithoutDeadHashesFormatted: core.PrettySize(raw.SharedSizeWithoutDeadHashes),

					TargetAvgPageSize:          raw.TargetSize.Avg(),
					TargetAvgPageSizeFormatted: core.PrettySize(raw.TargetSize.Avg()),
					TargetPageCount:            int(raw.TargetSize.Count),
				}
			},
		),
	}, nil
}
