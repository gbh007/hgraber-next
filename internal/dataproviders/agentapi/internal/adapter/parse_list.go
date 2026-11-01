package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) ParseList(ctx context.Context, u url.URL, onePage bool) (hproxymodel.List, error) {
	res, err := a.rawClient.APIV2ParsingListPost(ctx, &agentapi.APIV2ParsingListPostReq{
		URL:     u,
		OnePage: agentapi.NewOptBool(onePage),
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
