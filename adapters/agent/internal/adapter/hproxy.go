package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) HProxyList(ctx context.Context, u url.URL) (hproxymodel.List, error) {
	res, err := a.rawClient.APIHproxyParseListPost(ctx, &agentapi.APIHproxyParseListPostReq{
		URL: u,
	})
	if err != nil {
		return hproxymodel.List{}, fmt.Errorf("request: %w", err)
	}

	switch typedRes := res.(type) {
	case *agentapi.APIHproxyParseListPostOK:
		data := hproxymodel.List{}

		if typedRes.NextURL.Set {
			data.NextPage = &typedRes.NextURL.Value
		}

		for _, result := range typedRes.Results {
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

	case *agentapi.APIHproxyParseListPostBadRequest:
		return hproxymodel.List{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIHproxyParseListPostUnauthorized:
		return hproxymodel.List{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIHproxyParseListPostForbidden:
		return hproxymodel.List{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIHproxyParseListPostInternalServerError:
		return hproxymodel.List{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIInternalError, typedRes.Details.Value)

	default:
		return hproxymodel.List{}, agentmodel.ErrAgentAPIUnknownResponse
	}
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
		return hproxymodel.Book{}, fmt.Errorf("request: %w", err)
	}

	switch typedRes := res.(type) {
	case *agentapi.APIHproxyParseBookPostOK:
		var previewURL *url.URL

		if typedRes.PreviewURL.Set {
			previewURL = &typedRes.PreviewURL.Value
		}

		return hproxymodel.Book{
			Name:       typedRes.Name,
			ExURL:      typedRes.URL,
			PreviewURL: previewURL,
			PageCount:  typedRes.PageCount,
			Pages: pkg.Map(typedRes.Pages, func(p agentapi.APIHproxyParseBookPostOKPagesItem) hproxymodel.BookPage {
				return hproxymodel.BookPage{
					PageNumber:    p.PageNumber,
					ExtPreviewURL: p.URL,
				}
			}),
			Attributes: pkg.Map(
				typedRes.Attributes,
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

	case *agentapi.APIHproxyParseBookPostBadRequest:
		return hproxymodel.Book{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIHproxyParseBookPostUnauthorized:
		return hproxymodel.Book{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIHproxyParseBookPostForbidden:
		return hproxymodel.Book{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIHproxyParseBookPostInternalServerError:
		return hproxymodel.Book{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIInternalError, typedRes.Details.Value)

	default:
		return hproxymodel.Book{}, agentmodel.ErrAgentAPIUnknownResponse
	}
}
