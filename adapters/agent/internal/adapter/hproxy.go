package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) HProxyList(ctx context.Context, u url.URL) (hproxymodel.List, error) {
	res, err := a.rawClient.APIHproxyParseListPost(ctx, &agentapi.APIHproxyParseListPostReq{
		URL: u,
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
		case agentapi.APIHproxyParseListPostOKResultsItemTypeList:
			data.Pagination = append(data.Pagination, hproxymodel.ListPage{
				ExtURL: result.LinkURL,
				Name:   result.Name.Value,
			})

		case agentapi.APIHproxyParseListPostOKResultsItemTypeDetails:
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

	res, err := a.rawClient.APIHproxyParseBookPost(ctx, &agentapi.APIHproxyParseBookPostReq{
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
		Pages: pkg.Map(res.Pages, func(p agentapi.APIHproxyParseBookPostOKPagesItem) hproxymodel.BookPage {
			return hproxymodel.BookPage{
				PageNumber:    p.PageNumber,
				ExtPreviewURL: p.URL,
			}
		}),
		Attributes: pkg.Map(
			res.Attributes,
			func(attr agentapi.APIHproxyParseBookPostOKAttributesItem) hproxymodel.BookAttribute {
				return hproxymodel.BookAttribute{
					Code: attr.Code,
					Values: pkg.Map(
						attr.Values,
						func(
							v agentapi.APIHproxyParseBookPostOKAttributesItemValuesItem,
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
			}),
	}, nil
}
