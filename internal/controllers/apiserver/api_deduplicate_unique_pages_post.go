package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateUniquePagesPost(ctx context.Context, req *serverAPI.APIDeduplicateUniquePagesPostReq) (serverAPI.APIDeduplicateUniquePagesPostRes, error) {
	data, err := c.deduplicateUseCases.UniquePages(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateUniquePagesPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateUniquePagesPostOK{
		Pages: pkg.Map(data, func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		PagesWithoutDeadHashes: pkg.Map(pkg.SliceFilter(data, func(raw entities.PageWithDeadHash) bool {
			return !raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
		PagesOnlyDeadHashes: pkg.Map(pkg.SliceFilter(data, func(raw entities.PageWithDeadHash) bool {
			return raw.HasDeadHash
		}), func(raw entities.PageWithDeadHash) serverAPI.PageSimple {
			return c.convertSimplePage(raw.Page)
		}),
	}, nil
}
