package apiserver

import (
	"context"

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
		Pages: pkg.Map(data, c.convertSimplePage),
	}, nil
}
