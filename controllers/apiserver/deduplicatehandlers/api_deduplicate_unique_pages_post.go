package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateUniquePagesPost(ctx context.Context, req *serverAPI.APIDeduplicateUniquePagesPostReq) (serverAPI.APIDeduplicateUniquePagesPostRes, error) {
	data, err := c.deduplicateUseCases.UniquePages(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APIDeduplicateUniquePagesPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateUniquePagesPostOK{
		Pages: pkg.Map(data, c.apiCore.ConvertPreviewPage),
	}, nil
}
