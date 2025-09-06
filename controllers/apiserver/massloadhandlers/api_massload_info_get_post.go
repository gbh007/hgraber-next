package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *MassloadController) APIMassloadInfoGetPost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoGetPostReq,
) (serverapi.APIMassloadInfoGetPostRes, error) {
	ml, err := c.massloadUseCases.Massload(ctx, req.ID)
	if err != nil {
		return &serverapi.APIMassloadInfoGetPostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	resp := convertMassloadInfo(ml)

	return &resp, nil
}

func convertMassloadInfo(ml massloadmodel.Massload) serverapi.MassloadInfo {
	return serverapi.MassloadInfo{
		ID:          ml.ID,
		Name:        ml.Name,
		Description: apiservercore.OptString(ml.Description),
		Flags:       ml.Flags,

		PageSize:          apiservercore.OptInt64Pointer(ml.PageSize),
		PageSizeFormatted: apiservercore.OptString(core.PrettySizePointer(ml.PageSize)),
		PageCount:         apiservercore.OptInt64Pointer(ml.PageCount),

		FileSize:          apiservercore.OptInt64Pointer(ml.FileSize),
		FileSizeFormatted: apiservercore.OptString(core.PrettySizePointer(ml.FileSize)),
		FileCount:         apiservercore.OptInt64Pointer(ml.FileCount),

		BooksInSystem: apiservercore.OptInt64Pointer(ml.BookInSystem),

		BooksAhead:    apiservercore.OptInt64Pointer(ml.BooksAhead),
		NewBooks:      apiservercore.OptInt64Pointer(ml.NewBooks),
		ExistingBooks: apiservercore.OptInt64Pointer(ml.ExistingBooks),

		CreatedAt: ml.CreatedAt,
		UpdatedAt: apiservercore.OptTime(ml.UpdatedAt),
		ExternalLinks: pkg.Map(
			ml.ExternalLinks,
			func(link massloadmodel.ExternalLink) serverapi.MassloadInfoExternalLinksItem {
				return serverapi.MassloadInfoExternalLinksItem{
					URL:           link.URL,
					AutoCheck:     link.AutoCheck,
					BooksAhead:    apiservercore.OptInt64Pointer(link.BooksAhead),
					NewBooks:      apiservercore.OptInt64Pointer(link.NewBooks),
					ExistingBooks: apiservercore.OptInt64Pointer(link.ExistingBooks),
					CreatedAt:     link.CreatedAt,
					UpdatedAt:     apiservercore.OptTime(link.UpdatedAt),
				}
			},
		),
		Attributes: pkg.Map(ml.Attributes, func(attr massloadmodel.Attribute) serverapi.MassloadInfoAttributesItem {
			return serverapi.MassloadInfoAttributesItem{
				Code:  attr.Code,
				Value: attr.Value,

				PageSize:          apiservercore.OptInt64Pointer(attr.PageSize),
				PageSizeFormatted: apiservercore.OptString(core.PrettySizePointer(attr.PageSize)),
				PageCount:         apiservercore.OptInt64Pointer(attr.PageCount),

				FileSize:          apiservercore.OptInt64Pointer(attr.FileSize),
				FileSizeFormatted: apiservercore.OptString(core.PrettySizePointer(attr.FileSize)),
				FileCount:         apiservercore.OptInt64Pointer(attr.FileCount),

				BooksInSystem: apiservercore.OptInt64Pointer(attr.BookInSystem),

				CreatedAt: attr.CreatedAt,
				UpdatedAt: apiservercore.OptTime(attr.UpdatedAt),
			}
		}),
	}
}
