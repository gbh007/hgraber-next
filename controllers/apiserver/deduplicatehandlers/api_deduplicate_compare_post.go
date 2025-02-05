package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateComparePost(ctx context.Context, req *serverAPI.APIDeduplicateComparePostReq) (serverAPI.APIDeduplicateComparePostRes, error) {
	data, err := c.webAPIUseCases.BookCompare(ctx, req.OriginBookID, req.TargetBookID)
	if err != nil {
		return &serverAPI.APIDeduplicateComparePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateComparePostOK{
		Origin: c.apiCore.ConvertSimpleBook(data.OriginBook, data.OriginPreviewPage),
		Target: c.apiCore.ConvertSimpleBook(data.TargetBook, data.TargetPreviewPage),

		OriginPages: pkg.Map(data.OriginPages, c.apiCore.ConvertPreviewPage),
		BothPages:   pkg.Map(data.BothPages, c.apiCore.ConvertPreviewPage),
		TargetPages: pkg.Map(data.TargetPages, c.apiCore.ConvertPreviewPage),

		OriginAttributes: pkg.Map(data.OriginAttributes, apiservercore.ConvertBookAttribute),
		BothAttributes:   pkg.Map(data.BothAttributes, apiservercore.ConvertBookAttribute),
		TargetAttributes: pkg.Map(data.TargetAttributes, apiservercore.ConvertBookAttribute),

		OriginCoveredTarget:                  data.EntryPercentage,
		TargetCoveredOrigin:                  data.ReverseEntryPercentage,
		OriginCoveredTargetWithoutDeadHashes: data.EntryPercentageWithoutDeadHashes,
		TargetCoveredOriginWithoutDeadHashes: data.ReverseEntryPercentageWithoutDeadHashes,
	}, nil
}
