package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIBookRawPost(ctx context.Context, req *server.APIBookRawPostReq) (server.APIBookRawPostRes, error) {
	var (
		book      entities.BookFull
		err       error
		innerCode string
	)

	switch {
	case req.ID.IsSet():
		innerCode = WebAPIUseCaseCode
		book, err = c.webAPIUseCases.BookRaw(ctx, req.ID.Value)

	case req.URL.IsSet():
		innerCode = ParseUseCaseCode
		book, err = c.parseUseCases.BookByURL(ctx, req.URL.Value)

	default:
		return &server.APIBookRawPostBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("id and url is empty"),
		}, nil
	}

	if errors.Is(err, entities.BookNotFoundError) {
		return &server.APIBookRawPostNotFound{
			InnerCode: innerCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIBookRawPostInternalServerError{
			InnerCode: innerCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
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
