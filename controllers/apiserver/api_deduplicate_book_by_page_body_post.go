package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIDeduplicateBookByPageBodyPost(ctx context.Context, req *serverAPI.APIDeduplicateBookByPageBodyPostReq) (serverAPI.APIDeduplicateBookByPageBodyPostRes, error) {
	data, err := c.deduplicateUseCases.BookByPageEntryPercentage(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateBookByPageBodyPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateBookByPageBodyPostOK{
		Result: pkg.Map(data, func(raw bff.DeduplicateBookResult) serverAPI.APIDeduplicateBookByPageBodyPostOKResultItem {
			return serverAPI.APIDeduplicateBookByPageBodyPostOKResultItem{
				Book:                                 c.convertSimpleBook(raw.TargetBook, raw.PreviewPage),
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
		}),
	}, nil
}
