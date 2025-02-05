package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateUniquePagesPost(ctx context.Context, req *serverapi.APIDeduplicateUniquePagesPostReq) (serverapi.APIDeduplicateUniquePagesPostRes, error) {
	data, err := c.deduplicateUseCases.UniquePages(ctx, req.BookID)
	if err != nil {
		return &serverapi.APIDeduplicateUniquePagesPostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateUniquePagesPostOK{
		Pages: pkg.Map(data, c.apiCore.ConvertPreviewPage),
	}, nil
}
