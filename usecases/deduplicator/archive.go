package deduplicator

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"path"
	"slices"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/external"
)

func (uc *UseCase) ArchiveEntryPercentage(ctx context.Context, archiveBody io.Reader) ([]entities.DeduplicateArchiveResult, error) {
	bodyRaw, err := io.ReadAll(archiveBody)
	if err != nil {
		return nil, fmt.Errorf("read archive body: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(bodyRaw), int64(len(bodyRaw)))
	if err != nil {
		return nil, fmt.Errorf("create zip reader: %w", err)
	}

	archiveHashes := make([]entities.PageWithHash, 0, 30) // Точный размер заранее не определить
	md5Sums := make([]string, 0, 30)                      // Точный размер заранее не определить

	_, err = external.ReadArchive(ctx, zipReader, external.ReadArchiveOptions{
		HandlePageBody: func(ctx context.Context, pageNumber int, filename string, body io.Reader) error {
			hash, err := entities.HashFile(body)
			if err != nil {
				return fmt.Errorf("hash page (%d): %w", pageNumber, err)
			}

			md5Sums = append(md5Sums, hash.Md5Sum)
			archiveHashes = append(archiveHashes, entities.PageWithHash{
				Page: entities.Page{
					PageNumber: pageNumber,
					Ext:        path.Ext(filename),
					Downloaded: true,
				},
				FileHash: hash,
			})

			return nil
		},
		SkipInfo: true,
	})
	if err != nil {
		return nil, fmt.Errorf("read archive: %w", err)
	}

	bookIDs, err := uc.storage.BookIDsByMD5(ctx, md5Sums)
	if err != nil {
		return nil, fmt.Errorf("get books by md5 from storage: %w", err)
	}

	result := make([]entities.DeduplicateArchiveResult, 0, len(bookIDs))

	for _, bookID := range bookIDs {
		pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
		if err != nil {
			return nil, fmt.Errorf("get pages (%s) from storage: %w", bookID.String(), err)
		}

		bookShort, err := uc.storage.GetBook(ctx, bookID)
		if err != nil {
			return nil, fmt.Errorf("get book (%s) from storage: %w", bookID.String(), err)
		}

		result = append(result, entities.DeduplicateArchiveResult{
			TargetBookID:           bookID,
			OriginBookURL:          bookShort.OriginURL,
			EntryPercentage:        entities.EntryPercentageForPages(archiveHashes, pages, nil), // FIXME: а стоит ли игнорировать мертвые хеши?
			ReverseEntryPercentage: entities.EntryPercentageForPages(pages, archiveHashes, nil), // FIXME: а стоит ли игнорировать мертвые хеши?
		})
	}

	slices.SortFunc(result, func(a, b entities.DeduplicateArchiveResult) int {
		if a.EntryPercentage > b.EntryPercentage {
			return -1
		}

		if a.EntryPercentage < b.EntryPercentage {
			return 1
		}

		if a.ReverseEntryPercentage > b.ReverseEntryPercentage {
			return -1
		}

		if a.ReverseEntryPercentage < b.ReverseEntryPercentage {
			return 1
		}

		return 0
	})

	return result, nil
}
