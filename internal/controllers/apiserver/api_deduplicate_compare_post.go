package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateComparePost(ctx context.Context, req *serverAPI.APIDeduplicateComparePostReq) (serverAPI.APIDeduplicateComparePostRes, error) {
	data, err := c.deduplicateUseCases.BookPagesCompare(ctx, req.OriginBookID, req.TargetBookID)
	if err != nil {
		return &serverAPI.APIDeduplicateComparePostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateComparePostOK{
		Origin: c.convertSimpleBook(data.OriginBook, data.OriginPreviewPage),
		Target: c.convertSimpleBook(data.TargetBook, data.TargetPreviewPage),

		OriginPages: pkg.Map(data.OriginPages, func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		OriginPagesWithoutDeadHashes: pkg.Map(pkg.SliceFilter(data.OriginPages, func(raw entities.PageWithDeadHash) bool {
			return !raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		OriginPagesOnlyDeadHashes: pkg.Map(pkg.SliceFilter(data.OriginPages, func(raw entities.PageWithDeadHash) bool {
			return raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),

		BothPages: pkg.Map(data.BothPages, func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		BothPagesWithoutDeadHashes: pkg.Map(pkg.SliceFilter(data.BothPages, func(raw entities.PageWithDeadHash) bool {
			return !raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		BothPagesOnlyDeadHashes: pkg.Map(pkg.SliceFilter(data.BothPages, func(raw entities.PageWithDeadHash) bool {
			return raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),

		TargetPages: pkg.Map(data.TargetPages, func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		TargetPagesWithoutDeadHashes: pkg.Map(pkg.SliceFilter(data.TargetPages, func(raw entities.PageWithDeadHash) bool {
			return !raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		TargetPagesOnlyDeadHashes: pkg.Map(pkg.SliceFilter(data.TargetPages, func(raw entities.PageWithDeadHash) bool {
			return raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
	}, nil
}
