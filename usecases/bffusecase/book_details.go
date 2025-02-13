package bffusecase

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) BookDetails(ctx context.Context, bookID uuid.UUID) (bff.BookDetails, error) {
	book, err := uc.storage.GetBook(ctx, bookID)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get book: %w", err)
	}

	bookPages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get pages: %w", err)
	}

	fileCounts := make(map[core.FileHash]int, len(bookPages))

	md5Sums := make([]string, len(bookPages))
	for i, page := range bookPages {
		md5Sums[i] = page.Md5Sum

		fileCounts[page.FileHash] = 1 // Это условие (значение 1) нужно, чтобы дубликаты внутри книги не дали ложно-положительного срабатывания
	}

	pageDuplicates, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, md5Sums)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get page duplicates: %w", err)
	}

	for _, page := range pageDuplicates {
		if page.BookID == bookID {
			continue
		}

		if _, ok := fileCounts[page.FileHash]; ok {
			fileCounts[page.FileHash]++
		}
	}

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[core.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	result := bff.BookDetails{
		Book:  book,
		Pages: make([]bff.PreviewPage, 0, len(bookPages)),
	}

	fsInfos, err := uc.storage.FileStorages(ctx)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get file storages: %w", err)
	}

	fsDisposition := make(map[uuid.UUID]core.SizeWithCount, len(fsInfos))

	for _, page := range bookPages {
		_, hasDeadHash := existsDeadHashes[page.FileHash]

		preview := bff.PageWithHashToPreview(page)
		preview.HasDeadHash = bff.NewStatusFlag(hasDeadHash)

		result.Pages = append(result.Pages, preview)

		if preview.PageNumber == core.PageNumberForPreview {
			result.PreviewPage = preview
		}

		if c, ok := fileCounts[page.FileHash]; ok {
			if c > 1 {
				result.Size.Shared += page.Size
				result.Size.SharedCount++
			} else {
				result.Size.Unique += page.Size
				result.Size.UniqueCount++

				if !hasDeadHash {
					result.Size.UniqueWithoutDeadHashes += page.Size
					result.Size.UniqueWithoutDeadHashesCount++
				}
			}

			if hasDeadHash {
				result.Size.DeadHashes += page.Size
				result.Size.DeadHashesCount++
			}

			delete(fileCounts, page.FileHash) // Это нужно, чтобы дубликаты внутри книги не увеличивали уникальный объем
		} else {
			result.Size.InnerDuplicateCount++
		}

		result.Size.Total += page.Size

		fs := fsDisposition[page.FSID]
		fs.Count++
		fs.Size += page.Size
		fsDisposition[page.FSID] = fs
	}

	result.FSDisposition = make([]bff.BookDetailsFSDisposition, 0, len(fsDisposition))

	for _, fs := range fsInfos {
		v := fsDisposition[fs.ID]

		if v.Count > 0 {
			result.FSDisposition = append(result.FSDisposition, bff.BookDetailsFSDisposition{
				ID:            fs.ID,
				Name:          fs.Name,
				SizeWithCount: v,
			})
		}
	}

	slices.SortStableFunc(result.FSDisposition, func(a, b bff.BookDetailsFSDisposition) int {
		return int(b.Size) - int(a.Size)
	})

	attributes, err := uc.storage.BookAttributes(ctx, bookID)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get attributes: %w", err)
	}

	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return bff.BookDetails{}, fmt.Errorf("storage: get attributes info: %w", err)
	}

	result.Attributes = convertBookAttributes(
		convertAttributes(attributesInfo),
		attributes,
	)

	return result, nil
}

func convertAttributes(attributes []core.Attribute) map[string]core.Attribute {
	return pkg.SliceToMap(attributes, func(attribute core.Attribute) (string, core.Attribute) {
		return attribute.Code, attribute
	})
}

func convertBookAttributes(attributes map[string]core.Attribute, bookAttributes map[string][]string) []bff.AttributeToWeb {
	result := pkg.MapToSlice(bookAttributes, func(code string, values []string) bff.AttributeToWeb {
		return bff.AttributeToWeb{
			Code:   code,
			Name:   attributes[code].PluralName,
			Values: values,
		}
	})

	slices.SortStableFunc(result, func(a, b bff.AttributeToWeb) int {
		return attributes[a.Code].Order - attributes[b.Code].Order
	})

	return result
}
