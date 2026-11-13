package hproxyhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *HProxyHandlersController) APIHproxyBookPost(
	ctx context.Context,
	req *serverapi.APIHproxyBookPostReq,
) (*serverapi.APIHproxyBookPostOK, error) {
	var pageLimit *int

	if req.PageLimit.Set {
		pageLimit = &req.PageLimit.Value
	}

	book, err := c.hProxyUseCases.Book(ctx, req.URL, pageLimit)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIHproxyBookPostOK{
		Name:       book.Name,
		ExtURL:     book.ExURL,
		PreviewURL: c.filePreview(ctx, book.ExURL, book.PreviewURL),
		PageCount:  book.PageCount,
		ExistsIds:  book.ExistsIDs,
		Pages: lo.Map(book.Pages, func(p hproxymodel.BookPage, _ int) serverapi.APIHproxyBookPostOKPagesItem {
			return serverapi.APIHproxyBookPostOKPagesItem{
				PageNumber: p.PageNumber,
				PreviewURL: c.apiCore.GetHProxyFileURL(ctx, book.ExURL, p.ExtPreviewURL),
			}
		}),
		Attributes: lo.Map(
			book.Attributes,
			func(attr hproxymodel.BookAttribute, _ int) serverapi.APIHproxyBookPostOKAttributesItem {
				return serverapi.APIHproxyBookPostOKAttributesItem{
					Code: attr.Code,
					Name: attr.Name,
					Values: lo.Map(
						attr.Values,
						func(v hproxymodel.BookAttributeValue, _ int) serverapi.APIHproxyBookPostOKAttributesItemValuesItem {
							return serverapi.APIHproxyBookPostOKAttributesItemValuesItem{
								ExtName: v.ExtName,
								Name:    apiservercore.OptString(v.Name),
								ExtURL:  apiservercore.OptURL(v.ExtURL),
								MassloadsByName: lo.Map(
									v.MassloadsByName,
									func(ml hproxymodel.MassloadInfo, _ int) serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByNameItem { //nolint:lll // будет исправлено позднее
										return serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByNameItem{
											ID:   ml.ID,
											Name: ml.Name,
										}
									},
								),
								MassloadsByExtURL: lo.Map(
									v.MassloadsByURL,
									func(ml hproxymodel.MassloadInfo, _ int) serverapi.APIHproxyBookPostOKAttributesItemValuesItemMassloadsByExtURLItem { //nolint:lll // будет исправлено позднее
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
