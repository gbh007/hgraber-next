package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIBookRawPost(ctx context.Context, req *server.APIBookRawPostReq) (server.APIBookRawPostRes, error) {
	if !req.ID.IsSet() && !req.URL.IsSet() {
		return &server.APIBookRawPostBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("id and url is empty"),
		}, nil
	}

	var (
		book entities.BookFull
		err  error
	)

	switch {
	case req.ID.IsSet():
		book, err = c.webAPIUseCases.BookRaw(ctx, req.ID.Value)

		if errors.Is(err, entities.BookNotFoundError) {
			return &server.APIBookRawPostNotFound{
				InnerCode: WebAPIUseCaseCode,
				Details:   server.NewOptString(err.Error()),
			}, nil
		}

		if err != nil {
			return &server.APIBookRawPostInternalServerError{
				InnerCode: WebAPIUseCaseCode,
				Details:   server.NewOptString(err.Error()),
			}, nil
		}

	case req.URL.IsSet():
		book, err = c.parseUseCases.BookByURL(ctx, req.URL.Value)

		if errors.Is(err, entities.BookNotFoundError) {
			return &server.APIBookRawPostNotFound{
				InnerCode: ParseUseCaseCode,
				Details:   server.NewOptString(err.Error()),
			}, nil
		}

		if err != nil {
			return &server.APIBookRawPostInternalServerError{
				InnerCode: ParseUseCaseCode,
				Details:   server.NewOptString(err.Error()),
			}, nil
		}
	}

	return &server.BookRaw{
		ID:        book.Book.ID,
		CreateAt:  book.Book.CreateAt,
		OriginURL: optURL(book.Book.OriginURL),
		Name:      book.Book.Name,
		PageCount: book.Book.PageCount,
		Attributes: pkg.MapToSlice(book.Attributes, func(code string, values []string) server.BookRawAttributesItem {
			return server.BookRawAttributesItem{
				Code:   code,
				Values: values,
			}
		}),
		Pages: pkg.Map(book.Pages, func(p entities.Page) server.BookRawPagesItem {
			return server.BookRawPagesItem{
				PageNumber: p.PageNumber,
				OriginURL:  optURL(p.OriginURL),
				Ext:        p.Ext,
				CreateAt:   p.CreateAt,
				Downloaded: p.Downloaded,
				LoadAt:     optTime(p.LoadAt),
			}
		}),
		Labels: pkg.Map(book.Labels, func(l entities.BookLabel) server.BookRawLabelsItem {
			return server.BookRawLabelsItem{
				PageNumber: l.PageNumber,
				Name:       l.Name,
				Value:      l.Value,
				CreateAt:   l.CreateAt,
			}
		}),
	}, nil
}
