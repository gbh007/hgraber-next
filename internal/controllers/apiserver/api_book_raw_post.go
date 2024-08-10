package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookRawPost(ctx context.Context, req *serverAPI.APIBookRawPostReq) (serverAPI.APIBookRawPostRes, error) {
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
		return &serverAPI.APIBookRawPostBadRequest{
			InnerCode: ValidationCode,
			Details:   serverAPI.NewOptString("id and url is empty"),
		}, nil
	}

	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookRawPostNotFound{
			InnerCode: innerCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookRawPostInternalServerError{
			InnerCode: innerCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.BookRaw{
		ID:        book.Book.ID,
		CreateAt:  book.Book.CreateAt,
		OriginURL: optURL(book.Book.OriginURL),
		Name:      book.Book.Name,
		PageCount: book.Book.PageCount,
		Attributes: pkg.MapToSlice(book.Attributes, func(code string, values []string) serverAPI.BookRawAttributesItem {
			return serverAPI.BookRawAttributesItem{
				Code:   code,
				Values: values,
			}
		}),
		Pages: pkg.Map(book.Pages, func(p entities.Page) serverAPI.BookRawPagesItem {
			return serverAPI.BookRawPagesItem{
				PageNumber: p.PageNumber,
				OriginURL:  optURL(p.OriginURL),
				Ext:        p.Ext,
				CreateAt:   p.CreateAt,
				Downloaded: p.Downloaded,
				LoadAt:     optTime(p.LoadAt),
			}
		}),
		Labels: pkg.Map(book.Labels, func(l entities.BookLabel) serverAPI.BookRawLabelsItem {
			return serverAPI.BookRawLabelsItem{
				PageNumber: l.PageNumber,
				Name:       l.Name,
				Value:      l.Value,
				CreateAt:   l.CreateAt,
			}
		}),
	}, nil
}
