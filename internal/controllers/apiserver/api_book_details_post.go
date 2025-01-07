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
		ID:         book.Book.ID,
		Created:    book.Book.CreateAt,
		PreviewURL: c.getPagePreview(book.PreviewPage),
		Flags: serverAPI.BookFlags{
			ParsedName: book.Book.ParsedName(),
			ParsedPage: book.ParsedPages,
			IsVerified: book.Book.Verified,
			IsDeleted:  book.Book.Deleted,
			IsRebuild:  book.Book.IsRebuild,
		},
		Name:              book.Book.Name,
		PageCount:         book.Book.PageCount,
		PageLoadedPercent: book.PageDownloadPercent(),
		OriginURL:         optURL(book.Book.OriginURL),
		Attributes:        pkg.Map(book.Attributes, convertBookAttribute),
		Pages:             pkg.Map(book.Pages, c.convertSimplePageWithDeadHash),
		Size: serverAPI.OptBookDetailsSize{
			Value: serverAPI.BookDetailsSize{
				Unique:                           book.Size.Unique,
				UniqueWithoutDeadHashes:          book.Size.UniqueWithoutDeadHashes,
				Shared:                           book.Size.Shared,
				DeadHashes:                       book.Size.DeadHashes,
				Total:                            book.Size.Total,
				UniqueFormatted:                  entities.PrettySize(book.Size.Unique),
				UniqueWithoutDeadHashesFormatted: entities.PrettySize(book.Size.UniqueWithoutDeadHashes),
				SharedFormatted:                  entities.PrettySize(book.Size.Shared),
				DeadHashesFormatted:              entities.PrettySize(book.Size.DeadHashes),
				TotalFormatted:                   entities.PrettySize(book.Size.Total),
			},
			Set: book.Size.Total > 0,
		},
	}, nil
}
