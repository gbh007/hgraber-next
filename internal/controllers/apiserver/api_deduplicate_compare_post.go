package apiserver

import (
	"context"

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
		Origin:      c.convertSimpleBook(data.OriginBook, data.OriginPreviewPage),
		Target:      c.convertSimpleBook(data.TargetBook, data.TargetPreviewPage),
		OriginPages: pkg.Map(data.OriginPages, c.convertSimplePage),
		BothPages:   pkg.Map(data.BothPages, c.convertSimplePage),
		TargetPages: pkg.Map(data.TargetPages, c.convertSimplePage),
	}, nil
}
