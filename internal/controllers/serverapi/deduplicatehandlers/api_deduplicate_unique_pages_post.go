package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/bff"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
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
		Pages: lo.Map(data, func(page bff.PreviewPage, _ int) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
	}, nil
}
