package bookhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *BookHandlersController) APIBookListPost(
	ctx context.Context,
	req *serverapi.BookFilter,
) (serverapi.APIBookListPostRes, error) {
	filter := apiservercore.ConvertAPIBookFilter(*req)

	bookList, err := c.bffUseCases.BookList(ctx, filter)
	if err != nil {
		return &serverapi.APIBookListPostInternalServerError{
			InnerCode: apiservercore.BFFUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookListPostOK{
		Books: pkg.Map(bookList.Books, func(b bff.BookShort) serverapi.APIBookListPostOKBooksItem {
			return serverapi.APIBookListPostOKBooksItem{
				Info: c.apiCore.ConvertSimpleBook(ctx, b.Book, b.PreviewPage),
				Tags: b.Tags,
				ColorAttributes: pkg.Map(
					b.ColorAttributes,
					func(attr core.AttributeColor) serverapi.APIBookListPostOKBooksItemColorAttributesItem {
						return serverapi.APIBookListPostOKBooksItemColorAttributesItem{
							Code:            attr.Code,
							Value:           attr.Value,
							TextColor:       attr.TextColor,
							BackgroundColor: attr.BackgroundColor,
						}
					},
				),
			}
		}),
		Pages: pkg.Map(bookList.Pages, func(v int) serverapi.APIBookListPostOKPagesItem {
			return serverapi.APIBookListPostOKPagesItem{
				Value:       v,
				IsCurrent:   v == req.Pagination.Value.Page.Value,
				IsSeparator: v == -1,
			}
		}),
		Count: bookList.Count,
	}, nil
}
