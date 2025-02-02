package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIBookDetailsPost(ctx context.Context, req *serverAPI.APIBookDetailsPostReq) (serverAPI.APIBookDetailsPostRes, error) {
	book, err := c.bffUseCases.BookDetails(ctx, req.ID)
	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookDetailsPostNotFound{
			InnerCode: BFFUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookDetailsPostInternalServerError{
			InnerCode: BFFUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookDetailsPostOK{
		Info:              c.convertSimpleBook(book.Book, book.PreviewPage),
		PageLoadedPercent: book.PageDownloadPercent(),
		Attributes:        pkg.Map(book.Attributes, convertBookAttribute),
		Pages:             pkg.Map(book.Pages, c.convertPreviewPage),
		Size: serverAPI.OptAPIBookDetailsPostOKSize{
			Value: serverAPI.APIBookDetailsPostOKSize{
				Unique:                           book.Size.Unique,
				UniqueWithoutDeadHashes:          book.Size.UniqueWithoutDeadHashes,
				Shared:                           book.Size.Shared,
				DeadHashes:                       book.Size.DeadHashes,
				Total:                            book.Size.Total,
				UniqueFormatted:                  entities.PrettySize(book.Size.Unique),
				UniqueWithoutDeadHashesFormatted: entities.PrettySize(book.Size.UniqueWithoutDeadHashes),
				SharedFormatted:                  entities.PrettySize(book.Size.Shared),
				DeadHashesFormatted:              entities.PrettySize(book.Size.DeadHashes),
				TotalFormatted:                   entities.PrettySize(book.Size.Total),
				UniqueCount:                      book.Size.UniqueCount,
				UniqueWithoutDeadHashesCount:     book.Size.UniqueWithoutDeadHashesCount,
				SharedCount:                      book.Size.SharedCount,
				DeadHashesCount:                  book.Size.DeadHashesCount,
				InnerDuplicateCount:              book.Size.InnerDuplicateCount,
				AvgPageSize:                      book.AvgPageSize(),
				AvgPageSizeFormatted:             entities.PrettySize(book.AvgPageSize()),
			},
			Set: book.Size.Total > 0,
		},
		FsDisposition: pkg.Map(book.FSDisposition, func(raw entities.BFFBookDetailsFSDisposition) serverAPI.APIBookDetailsPostOKFsDispositionItem {
			return serverAPI.APIBookDetailsPostOKFsDispositionItem{
				ID:   raw.ID,
				Name: raw.Name,
				Files: serverAPI.FSDBFilesInfo{
					Count:         raw.Count,
					Size:          raw.Size,
					SizeFormatted: entities.PrettySize(raw.Size),
				},
			}
		}),
	}, nil
}
