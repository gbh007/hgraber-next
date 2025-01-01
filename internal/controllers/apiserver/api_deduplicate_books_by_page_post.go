package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIDeduplicateBooksByPagePost(ctx context.Context, req *serverAPI.APIDeduplicateBooksByPagePostReq) (serverAPI.APIDeduplicateBooksByPagePostRes, error) {
	data, err := c.deduplicateUseCases.BooksByPage(ctx, req.BookID, req.PageNumber)
	if err != nil {
		return &serverAPI.APIDeduplicateBooksByPagePostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIDeduplicateBooksByPagePostOK{
		Books: pkg.Map(data, func(raw entities.BookWithPreviewPage) serverAPI.BookSimple {
			return c.convertSimpleBook(raw.Book, raw.PreviewPage)
		}),
	}, nil
}
