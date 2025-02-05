package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIBookDetailsPost(ctx context.Context, req *serverAPI.APIBookDetailsPostReq) (serverAPI.APIBookDetailsPostRes, error) {
	book, err := c.bffUseCases.BookDetails(ctx, req.ID)
	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookDetailsPostNotFound{
			InnerCode: apiservercore.BFFUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookDetailsPostInternalServerError{
			InnerCode: apiservercore.BFFUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookDetailsPostOK{
		Info:              c.apiCore.ConvertSimpleBook(book.Book, book.PreviewPage),
		PageLoadedPercent: book.PageDownloadPercent(),
		Attributes:        pkg.Map(book.Attributes, apiservercore.ConvertBookAttribute),
		Pages:             pkg.Map(book.Pages, c.apiCore.ConvertPreviewPage),
		Size: serverAPI.OptAPIBookDetailsPostOKSize{
			Value: serverAPI.APIBookDetailsPostOKSize{
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
		FsDisposition: pkg.Map(book.FSDisposition, func(raw bff.BookDetailsFSDisposition) serverAPI.APIBookDetailsPostOKFsDispositionItem {
			return serverAPI.APIBookDetailsPostOKFsDispositionItem{
				ID:   raw.ID,
				Name: raw.Name,
				Files: serverAPI.FSDBFilesInfo{
					Count:         raw.Count,
					Size:          raw.Size,
					SizeFormatted: core.PrettySize(raw.Size),
				},
			}
		}),
	}, nil
}
