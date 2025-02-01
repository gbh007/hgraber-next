package apiserver

import (
	"context"

	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateComparePost(ctx context.Context, req *serverAPI.APIDeduplicateComparePostReq) (serverAPI.APIDeduplicateComparePostRes, error) {
	data, err := c.webAPIUseCases.BookCompare(ctx, req.OriginBookID, req.TargetBookID)
	if err != nil {
		return &serverAPI.APIDeduplicateComparePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateComparePostOK{
		Origin: c.convertSimpleBook(data.OriginBook, data.OriginPreviewPage),
		Target: c.convertSimpleBook(data.TargetBook, data.TargetPreviewPage),

		OriginPages: pkg.Map(data.OriginPages, c.convertPreviewPage),
		BothPages:   pkg.Map(data.BothPages, c.convertPreviewPage),
		TargetPages: pkg.Map(data.TargetPages, c.convertPreviewPage),

		OriginAttributes: pkg.Map(data.OriginAttributes, convertBookAttribute),
		BothAttributes:   pkg.Map(data.BothAttributes, convertBookAttribute),
		TargetAttributes: pkg.Map(data.TargetAttributes, convertBookAttribute),

		OriginCoveredTarget:                  data.EntryPercentage,
		TargetCoveredOrigin:                  data.ReverseEntryPercentage,
		OriginCoveredTargetWithoutDeadHashes: data.EntryPercentageWithoutDeadHashes,
		TargetCoveredOriginWithoutDeadHashes: data.ReverseEntryPercentageWithoutDeadHashes,
	}, nil
}
