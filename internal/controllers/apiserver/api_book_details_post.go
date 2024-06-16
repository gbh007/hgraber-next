package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIBookDetailsPost(ctx context.Context, req *server.APIBookDetailsPostReq) (server.APIBookDetailsPostRes, error) {
	book, err := c.webAPIUseCases.Book(ctx, req.ID)
	if errors.Is(err, entities.BookNotFoundError) {
		return &server.APIBookDetailsPostNotFound{
			InnerCode: WebAPIUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIBookDetailsPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	previewURL := server.OptURI{}

	if book.PreviewPage.Downloaded {
		previewURL = server.NewOptURI(c.getFileURL(
			book.PreviewPage.FileID,
			book.PreviewPage.Ext,
		))
	}

	return &server.BookDetails{
		ID:                book.Book.ID,
		Created:           book.Book.CreateAt,
		PreviewURL:        previewURL,
		ParsedName:        book.ParsedName(),
		Name:              book.Book.Name,
		ParsedPage:        book.ParsedPages,
		PageCount:         book.Book.PageCount,
		PageLoadedPercent: book.PageDownloadPercent(),
		Attributes: pkg.Map(book.Attributes, func(a entities.AttributeToWeb) server.BookDetailsAttributesItem {
			return server.BookDetailsAttributesItem{
				Name:   a.Name,
				Values: a.Values,
			}
		}),
		Pages: pkg.Map(book.Pages, func(p entities.Page) server.BookDetailsPagesItem {
			previewURL := server.OptURI{}

			if p.Downloaded {
				previewURL = server.NewOptURI(c.getFileURL(
					p.FileID,
					p.Ext,
				))
			}

			return server.BookDetailsPagesItem{
				PageNumber: p.PageNumber,
				PreviewURL: previewURL,
			}
		}),
	}, nil
}
