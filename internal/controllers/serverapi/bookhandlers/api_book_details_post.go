package bookhandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/bff"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *BookHandlersController) APIBookDetailsPost(
	ctx context.Context,
	req *serverapi.APIBookDetailsPostReq,
) (*serverapi.APIBookDetailsPostOK, error) {
	book, err := c.bffUseCases.BookDetails(ctx, req.ID)
	if errors.Is(err, core.ErrBookNotFound) {
		return nil, apiservercore.APIError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	//nolint:staticcheck // поля используются в API
	return &serverapi.APIBookDetailsPostOK{
		Info:              c.apiCore.ConvertSimpleBook(ctx, book.Book, book.PreviewPage),
		PageLoadedPercent: book.PageDownloadPercent(),
		Attributes: lo.Map(book.Attributes, func(attr bff.AttributeToWeb, _ int) serverapi.BookAttribute {
			return apiservercore.ConvertBookAttribute(attr)
		}),
		Pages: lo.Map(book.Pages, func(page bff.PreviewPage, _ int) serverapi.PageSimple {
			return c.apiCore.ConvertPreviewPage(ctx, page)
		}),
		Size: serverapi.OptAPIBookDetailsPostOKSize{
			Value: serverapi.APIBookDetailsPostOKSize{
				Unique:                           book.Size.Unique,
				UniqueWithoutDeadHashes:          book.Size.UniqueWithoutDeadHashes,
				Shared:                           book.Size.Shared,
				DeadHashes:                       book.Size.DeadHashes,
				Total:                            book.Size.Total,
				UniqueFormatted:                  core.PrettySize(book.Size.Unique),
				UniqueWithoutDeadHashesFormatted: core.PrettySize(book.Size.UniqueWithoutDeadHashes),
				SharedFormatted:                  core.PrettySize(book.Size.Shared),
				DeadHashesFormatted:              core.PrettySize(book.Size.DeadHashes),
				TotalFormatted:                   core.PrettySize(book.Size.Total),
				UniqueCount:                      book.Size.UniqueCount,
				UniqueWithoutDeadHashesCount:     book.Size.UniqueWithoutDeadHashesCount,
				SharedCount:                      book.Size.SharedCount,
				DeadHashesCount:                  book.Size.DeadHashesCount,
				InnerDuplicateCount:              book.Size.InnerDuplicateCount,
				AvgPageSize:                      book.AvgPageSize(),
				AvgPageSizeFormatted:             core.PrettySize(book.AvgPageSize()),
			},
			Set: book.Size.Total > 0,
		},
		FsDisposition: lo.Map(
			book.FSDisposition,
			func(raw bff.BookDetailsFSDisposition, _ int) serverapi.APIBookDetailsPostOKFsDispositionItem {
				return serverapi.APIBookDetailsPostOKFsDispositionItem{
					ID:   raw.ID,
					Name: raw.Name,
					Files: serverapi.FSDBFilesInfo{
						Count:         raw.Count,
						Size:          raw.Size,
						SizeFormatted: core.PrettySize(raw.Size),
					},
				}
			},
		),
	}, nil
}
