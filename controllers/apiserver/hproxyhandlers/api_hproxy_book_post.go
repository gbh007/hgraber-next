package hproxyhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *HProxyHandlersController) APIHproxyBookPost(
	ctx context.Context,
	req *serverapi.APIHproxyBookPostReq,
) (serverapi.APIHproxyBookPostRes, error) {
	var pageLimit *int

	if req.PageLimit.Set {
		pageLimit = &req.PageLimit.Value
	}

	book, err := c.hProxyUseCases.Book(ctx, req.URL, pageLimit)
	if err != nil {
		return &serverapi.APIHproxyBookPostInternalServerError{
			InnerCode: apiservercore.HProxyUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIHproxyBookPostOK{
		Name:       book.Name,
		ExtURL:     book.ExURL,
		PreviewURL: c.filePreview(book.ExURL, book.PreviewURL),
		PageCount:  book.PageCount,
		ExistsIds:  book.ExistsIDs,
		Pages: pkg.Map(book.Pages, func(p hproxymodel.BookPage) serverapi.APIHproxyBookPostOKPagesItem {
			return serverapi.APIHproxyBookPostOKPagesItem{
				PageNumber: p.PageNumber,
				PreviewURL: c.apiCore.GetHProxyFileURL(book.ExURL, p.ExtPreviewURL),
			}
		}),
		Attributes: pkg.Map(
			book.Attributes,
			func(attr hproxymodel.BookAttribute) serverapi.APIHproxyBookPostOKAttributesItem {
				return serverapi.APIHproxyBookPostOKAttributesItem{
					Code: attr.Code,
					Name: attr.Name,
					Values: pkg.Map(
						attr.Values,
						func(v hproxymodel.BookAttributeValue) serverapi.APIHproxyBookPostOKAttributesItemValuesItem {
							return serverapi.APIHproxyBookPostOKAttributesItemValuesItem{
								ExtName: v.ExtName,
								Name:    apiservercore.OptString(v.Name),
								ExtURL:  apiservercore.OptURL(v.ExtURL),
								MassloadsByName: pkg.Map(
									v.MassloadsByName,
									func(ml hproxymodel.MassloadInfo) serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByNameItem { //nolint:lll // будет исправлено позднее
										return serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByNameItem{
											ID:   ml.ID,
											Name: ml.Name,
										}
									},
								),
								MassloadsByExtURL: pkg.Map(
									v.MassloadsByURL,
									func(ml hproxymodel.MassloadInfo) serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByExtURLItem { //nolint:lll // будет исправлено позднее
										return serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByExtURLItem{ //nolint:lll // будет исправлено позднее
											ID:   ml.ID,
											Name: ml.Name,
										}
									},
								),
							}
						},
					),
				}
			},
		),
	}, nil
}
