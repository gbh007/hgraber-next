package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/bff"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *DeduplicateHandlersController) APIDeduplicateBooksByPagePost(
	ctx context.Context,
	req *serverapi.APIDeduplicateBooksByPagePostReq,
) (*serverapi.APIDeduplicateBooksByPagePostOK, error) {
	data, err := c.deduplicateUseCases.BooksByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIDeduplicateBooksByPagePostOK{
		Books: pkg.Map(data, func(raw bff.BookWithPreviewPage) serverapi.BookSimple {
			return c.apiCore.ConvertSimpleBook(ctx, raw.Book, raw.PreviewPage)
		}),
	}, nil
}
