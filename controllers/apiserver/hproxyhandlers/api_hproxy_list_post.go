package hproxyhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *HProxyHandlersController) APIHproxyListPost(
	ctx context.Context,
	req *serverapi.APIHproxyListPostReq,
) (*serverapi.APIHproxyListPostOK, error) {
	data, err := c.hProxyUseCases.List(ctx, req.URL)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.HProxyUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIHproxyListPostOK{
		Books: pkg.Map(data.Books, func(b hproxymodel.ListBook) serverapi.APIHproxyListPostOKBooksItem {
			return serverapi.APIHproxyListPostOKBooksItem{
				ExtURL:     b.ExtURL,
				Name:       apiservercore.OptString(b.Name),
				PreviewURL: c.filePreview(b.ExtURL, b.ExtPreviewURL),
				ExistsIds:  b.ExistsIDs,
			}
		}),
		Pagination: pkg.Map(data.Pagination, func(p hproxymodel.ListPage) serverapi.APIHproxyListPostOKPaginationItem {
			return serverapi.APIHproxyListPostOKPaginationItem{
				ExtURL: p.ExtURL,
				Name:   p.Name,
			}
		}),
	}, nil
}
