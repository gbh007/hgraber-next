package hproxyhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *HProxyHandlersController) APIHproxyListPost(ctx context.Context, req *serverapi.APIHproxyListPostReq) (serverapi.APIHproxyListPostRes, error) {
	data, err := c.hProxyUseCases.List(ctx, req.URL)
	if err != nil {
		return &serverapi.APIHproxyListPostInternalServerError{
			InnerCode: apiservercore.HProxyUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIHproxyListPostOK{
		Books: pkg.Map(data.Books, func(b hproxymodel.ListBook) serverapi.APIHproxyListPostOKBooksItem {
			return serverapi.APIHproxyListPostOKBooksItem{
				ExtURL:     b.ExtURL,
				Name:       apiservercore.OptString(b.Name),
				PreviewURL: c.filePreview(b.ExtURL, b.ExtPreviewURL),
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
