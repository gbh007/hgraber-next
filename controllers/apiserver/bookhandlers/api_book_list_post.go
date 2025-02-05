package bookhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *BookHandlersController) APIBookListPost(ctx context.Context, req *serverAPI.BookFilter) (serverAPI.APIBookListPostRes, error) {
	filter := apiservercore.ConvertAPIBookFilter(*req)

	bookList, err := c.bffUseCases.BookList(ctx, filter)
	if err != nil {
		return &serverAPI.APIBookListPostInternalServerError{
			InnerCode: apiservercore.BFFUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookListPostOK{
		Books: pkg.Map(bookList.Books, func(b bff.BookShort) serverAPI.APIBookListPostOKBooksItem {
			return serverAPI.APIBookListPostOKBooksItem{
				Info: c.apiCore.ConvertSimpleBook(b.Book, b.PreviewPage),
				Tags: b.Tags,
			}
		}),
		Pages: pkg.Map(bookList.Pages, func(v int) serverAPI.APIBookListPostOKPagesItem {
			return serverAPI.APIBookListPostOKPagesItem{
				Value:       v,
				IsCurrent:   v == req.Pagination.Value.Page.Value,
				IsSeparator: v == -1,
			}
		}),
		Count: bookList.Count,
	}, nil
}
