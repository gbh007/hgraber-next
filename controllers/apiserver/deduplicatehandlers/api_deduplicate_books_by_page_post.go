package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateBooksByPagePost(
	ctx context.Context,
	req *serverapi.APIDeduplicateBooksByPagePostReq,
) (serverapi.APIDeduplicateBooksByPagePostRes, error) {
	data, err := c.deduplicateUseCases.BooksByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverapi.APIDeduplicateBooksByPagePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIDeduplicateBooksByPagePostOK{
		Books: pkg.Map(data, func(raw bff.BookWithPreviewPage) serverapi.BookSimple {
			return c.apiCore.ConvertSimpleBook(ctx, raw.Book, raw.PreviewPage)
		}),
	}, nil
}
