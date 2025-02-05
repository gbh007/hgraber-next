package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIDeduplicateBooksByPagePost(ctx context.Context, req *serverAPI.APIDeduplicateBooksByPagePostReq) (serverAPI.APIDeduplicateBooksByPagePostRes, error) {
	data, err := c.deduplicateUseCases.BooksByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverAPI.APIDeduplicateBooksByPagePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateBooksByPagePostOK{
		Books: pkg.Map(data, func(raw bff.BookWithPreviewPage) serverAPI.BookSimple {
			return c.apiCore.ConvertSimpleBook(raw.Book, raw.PreviewPage)
		}),
	}, nil
}
