package hproxyhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *HProxyHandlersController) APIHproxyListPost(
	ctx context.Context,
	req *serverapi.APIHproxyListPostReq,
) (*serverapi.APIHproxyListPostOK, error) {
	data, err := c.hProxyUseCases.List(ctx, req.URL)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIHproxyListPostOK{
		MassloadIds: lo.Map(data.Massloads, func(ml hproxymodel.MassloadInfo, _ int) int { return ml.ID }),
		Books: lo.Map(data.Books, func(b hproxymodel.ListBook, _ int) serverapi.APIHproxyListPostOKBooksItem {
			return serverapi.APIHproxyListPostOKBooksItem{
				ExtURL:     b.ExtURL,
				Name:       apiservercore.OptString(b.Name),
				PreviewURL: c.filePreview(ctx, b.ExtURL, b.ExtPreviewURL),
				ExistsIds:  b.ExistsIDs,
			}
		}),
		Pagination: lo.Map(data.Pagination, func(p hproxymodel.ListPage, _ int) serverapi.APIHproxyListPostOKPaginationItem {
			return serverapi.APIHproxyListPostOKPaginationItem{
				ExtURL: p.ExtURL,
				Name:   p.Name,
			}
		}),
	}, nil
}
