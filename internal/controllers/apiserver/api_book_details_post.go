package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookDetailsPost(ctx context.Context, req *serverAPI.APIBookDetailsPostReq) (serverAPI.APIBookDetailsPostRes, error) {
	book, err := c.webAPIUseCases.Book(ctx, req.ID)
	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookDetailsPostNotFound{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookDetailsPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.BookDetails{
		ID:                book.Book.ID,
		Created:           book.Book.CreateAt,
		PreviewURL:        c.getPagePreview(book.PreviewPage),
		ParsedName:        book.ParsedName(),
		Name:              book.Book.Name,
		ParsedPage:        book.ParsedPages,
		PageCount:         book.Book.PageCount,
		PageLoadedPercent: book.PageDownloadPercent(),
		Attributes: pkg.Map(book.Attributes, func(a entities.AttributeToWeb) serverAPI.BookDetailsAttributesItem {
			return serverAPI.BookDetailsAttributesItem{
				Name:   a.Name,
				Values: a.Values,
			}
		}),
		Pages: pkg.Map(book.Pages, func(p entities.Page) serverAPI.BookDetailsPagesItem {
			return serverAPI.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				PreviewURL: c.getPagePreview(p),
			}
		}),
	}, nil
}
