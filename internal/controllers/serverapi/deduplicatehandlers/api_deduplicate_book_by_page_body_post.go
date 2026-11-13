package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/bff"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *DeduplicateHandlersController) APIDeduplicateBookByPageBodyPost(
	ctx context.Context,
	req *serverapi.APIDeduplicateBookByPageBodyPostReq,
) (*serverapi.APIDeduplicateBookByPageBodyPostOK, error) {
	data, err := c.deduplicateUseCases.BookByPageEntryPercentage(ctx, req.BookID)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	//nolint:staticcheck // поля используются в API
	return &serverapi.APIDeduplicateBookByPageBodyPostOK{
		Result: lo.Map(
			data,
			func(raw bff.DeduplicateBookResult, _ int) serverapi.APIDeduplicateBookByPageBodyPostOKResultItem {
				//nolint:golines,lll // будет исправлено позднее
				return serverapi.APIDeduplicateBookByPageBodyPostOKResultItem{
					Book:                                 c.apiCore.ConvertSimpleBook(ctx, raw.TargetBook, raw.PreviewPage),
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
