package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/bff"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *DeduplicateHandlersController) APIDeduplicateComparePost(
	ctx context.Context,
	req *serverapi.APIDeduplicateComparePostReq,
) (*serverapi.APIDeduplicateComparePostOK, error) {
	data, err := c.bffUseCases.BookCompare(ctx, req.OriginBookID, req.TargetBookID)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	//nolint:staticcheck // поля используются в API
	return &serverapi.APIDeduplicateComparePostOK{
		Origin: c.apiCore.ConvertSimpleBook(ctx, data.OriginBook, data.OriginPreviewPage),
		Target: c.apiCore.ConvertSimpleBook(ctx, data.TargetBook, data.TargetPreviewPage),

		OriginPages: lo.Map(data.OriginPages, func(page bff.PreviewPage, _ int) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
		BothPages: lo.Map(data.BothPages, func(page bff.PreviewPage, _ int) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
		TargetPages: lo.Map(data.TargetPages, func(page bff.PreviewPage, _ int) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),

		OriginAttributes: lo.Map(data.OriginAttributes, func(attr bff.AttributeToWeb, _ int) serverapi.BookAttribute {
			return apiservercore.ConvertBookAttribute(attr)
		}),
		BothAttributes: lo.Map(data.BothAttributes, func(attr bff.AttributeToWeb, _ int) serverapi.BookAttribute {
			return apiservercore.ConvertBookAttribute(attr)
		}),
		TargetAttributes: lo.Map(data.TargetAttributes, func(attr bff.AttributeToWeb, _ int) serverapi.BookAttribute {
			return apiservercore.ConvertBookAttribute(attr)
		}),

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
