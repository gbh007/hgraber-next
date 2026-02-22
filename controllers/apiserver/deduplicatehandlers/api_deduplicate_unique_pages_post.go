package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateUniquePagesPost(
	ctx context.Context,
	req *serverapi.APIDeduplicateUniquePagesPostReq,
) (*serverapi.APIDeduplicateUniquePagesPostOK, error) {
	data, err := c.deduplicateUseCases.UniquePages(ctx, req.BookID)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIDeduplicateUniquePagesPostOK{
		Pages: pkg.Map(data, func(page bff.PreviewPage) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
	}, nil
}
