package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/samber/lo"
)

func (a *Adapter) HProxyList(ctx context.Context, u url.URL) (hproxymodel.List, error) {
	res, err := a.rawClient.APIV2ParsingListPost(ctx, &agentapi.APIV2ParsingListPostReq{
		URL:     u,
		OnePage: agentapi.NewOptBool(true),
	})
	if err != nil {
		return hproxymodel.List{}, enrichError(err)
	}

	data := hproxymodel.List{}

	if res.NextURL.Set {
		data.NextPage = &res.NextURL.Value
	}

	for _, result := range res.Results {
		switch result.Type {
		case agentapi.APIV2ParsingListPostOKResultsItemTypeList:
			data.Pagination = append(data.Pagination, hproxymodel.ListPage{
				ExtURL: result.LinkURL,
				Name:   result.Name.Value,
			})

		case agentapi.APIV2ParsingListPostOKResultsItemTypeDetails:
			var previewURL *url.URL

			if result.PreviewURL.Set {
				previewURL = &result.PreviewURL.Value
			}

			data.Books = append(data.Books, hproxymodel.ListBook{
				ExtURL:        result.LinkURL,
				Name:          result.Name.Value,
				ExtPreviewURL: previewURL,
			})
		}
	}

	return data, nil
}

func (a *Adapter) HProxyBook(ctx context.Context, u url.URL, pageLimit *int) (hproxymodel.Book, error) {
	pl := agentapi.OptInt{}

	if pageLimit != nil {
		pl = agentapi.NewOptInt(*pageLimit)
	}

	res, err := a.rawClient.APIV2ParsingBookPost(ctx, &agentapi.APIV2ParsingBookPostReq{
		URL:       u,
		PageLimit: pl,
	})
	if err != nil {
		return hproxymodel.Book{}, enrichError(err)
	}

	var previewURL *url.URL

	if res.PreviewURL.Set {
		previewURL = &res.PreviewURL.Value
	}

	return hproxymodel.Book{
		Name:       res.Name,
		ExURL:      res.URL,
		PreviewURL: previewURL,
		PageCount:  res.PageCount,
		Pages: lo.Map(res.Pages, func(p agentapi.BookDetailsPagesItem, _ int) hproxymodel.BookPage {
			return hproxymodel.BookPage{
				PageNumber:    p.PageNumber,
				ExtPreviewURL: p.URL,
			}
		}),
		Attributes: lo.Map(
			res.Attributes,
			func(attr agentapi.BookDetailsAttributesItem, _ int) hproxymodel.BookAttribute {
				return hproxymodel.BookAttribute{
					Code: attr.Code,
					Values: lo.Map(
						attr.Values,
						func(
							v agentapi.BookDetailsAttributesItemValuesItem,
							_ int,
						) hproxymodel.BookAttributeValue {
							var u *url.URL

							if v.URL.Set {
								u = &v.URL.Value
							}

							return hproxymodel.BookAttributeValue{
								ExtName: v.Name,
								ExtURL:  u,
							}
						},
					),
				}
			},
		),
	}, nil
}
