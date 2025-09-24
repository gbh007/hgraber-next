package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateComparePost(
	ctx context.Context,
	req *serverapi.APIDeduplicateComparePostReq,
) (serverapi.APIDeduplicateComparePostRes, error) {
	data, err := c.bffUseCases.BookCompare(ctx, req.OriginBookID, req.TargetBookID)
	if err != nil {
		return &serverapi.APIDeduplicateComparePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateComparePostOK{
		Origin: c.apiCore.ConvertSimpleBook(ctx, data.OriginBook, data.OriginPreviewPage),
		Target: c.apiCore.ConvertSimpleBook(ctx, data.TargetBook, data.TargetPreviewPage),

		OriginPages: pkg.Map(data.OriginPages, func(page bff.PreviewPage) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
		BothPages: pkg.Map(data.BothPages, func(page bff.PreviewPage) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
		TargetPages: pkg.Map(data.TargetPages, func(page bff.PreviewPage) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),

		OriginAttributes: pkg.Map(data.OriginAttributes, apiservercore.ConvertBookAttribute),
		BothAttributes:   pkg.Map(data.BothAttributes, apiservercore.ConvertBookAttribute),
		TargetAttributes: pkg.Map(data.TargetAttributes, apiservercore.ConvertBookAttribute),

		OriginCoveredTarget:                  data.EntryPercentage,
		TargetCoveredOrigin:                  data.ReverseEntryPercentage,
		OriginCoveredTargetWithoutDeadHashes: data.EntryPercentageWithoutDeadHashes,
		TargetCoveredOriginWithoutDeadHashes: data.ReverseEntryPercentageWithoutDeadHashes,

		OriginBookSize:             serverapi.NewOptInt64(data.OriginSize.Size),
		OriginBookSizeFormatted:    serverapi.NewOptString(core.PrettySize(data.OriginSize.Size)),
		OriginPageAvgSize:          serverapi.NewOptInt64(data.OriginSize.Avg()),
		OriginPageAvgSizeFormatted: serverapi.NewOptString(core.PrettySize(data.OriginSize.Avg())),

		TargetBookSize:             serverapi.NewOptInt64(data.TargetSize.Size),
		TargetBookSizeFormatted:    serverapi.NewOptString(core.PrettySize(data.TargetSize.Size)),
		TargetPageAvgSize:          serverapi.NewOptInt64(data.TargetSize.Avg()),
		TargetPageAvgSizeFormatted: serverapi.NewOptString(core.PrettySize(data.TargetSize.Avg())),
	}, nil
}
