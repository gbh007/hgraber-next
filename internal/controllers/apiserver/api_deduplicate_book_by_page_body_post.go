package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
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
		Result: pkg.Map(data, func(raw entities.DeduplicateBookResult) serverAPI.APIDeduplicateBookByPageBodyPostOKResultItem {
			return serverAPI.APIDeduplicateBookByPageBodyPostOKResultItem{
				Book:                                 c.convertSimpleBook(raw.TargetBook, raw.PreviewPage),
				OriginCoveredTarget:                  raw.EntryPercentage,
				TargetCoveredOrigin:                  raw.ReverseEntryPercentage,
				OriginCoveredTargetWithoutDeadHashes: raw.EntryPercentageWithoutDeadHashes,
				TargetCoveredOriginWithoutDeadHashes: raw.ReverseEntryPercentageWithoutDeadHashes,
			}
		}),
	}, nil
}
