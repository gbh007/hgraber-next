package bff

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) BookDetails(ctx context.Context, bookID uuid.UUID) (entities.BFFBookDetails, error) {
	book, err := uc.storage.GetBook(ctx, bookID)
	if err != nil {
		return entities.BFFBookDetails{}, fmt.Errorf("storage: get book: %w", err)
	}

	bookPages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return entities.BFFBookDetails{}, fmt.Errorf("storage: get pages: %w", err)
	}

	fileCounts := make(map[entities.FileHash]int, len(bookPages))

	md5Sums := make([]string, len(bookPages))
	for i, page := range bookPages {
		md5Sums[i] = page.Md5Sum

		fileCounts[page.FileHash] = 1 // Это условие (значение 1) нужно, чтобы дубликаты внутри книги не дали ложно-положительного срабатывания
	}

	pageDuplicates, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, md5Sums)
	if err != nil {
		return entities.BFFBookDetails{}, fmt.Errorf("storage: get page duplicates: %w", err)
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
		return entities.BFFBookDetails{}, fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[entities.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	result := entities.BFFBookDetails{
		Book:  book,
		Pages: make([]entities.PreviewPage, 0, len(bookPages)),
	}

	for _, page := range bookPages {
		_, hasDeadHash := existsDeadHashes[page.FileHash]

		preview := page.ToPreview()
		preview.HasDeadHash = &hasDeadHash

		result.Pages = append(result.Pages, preview)

		if preview.PageNumber == entities.PageNumberForPreview {
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
	}

	attributes, err := uc.storage.BookAttributes(ctx, bookID)
	if err != nil {
		return entities.BFFBookDetails{}, fmt.Errorf("storage: get attributes: %w", err)
	}

	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return entities.BFFBookDetails{}, fmt.Errorf("storage: get attributes info: %w", err)
	}

	result.Attributes = convertBookAttributes(
		convertAttributes(attributesInfo),
		attributes,
	)

	return result, nil
}

func convertAttributes(attributes []entities.Attribute) map[string]entities.Attribute {
	return pkg.SliceToMap(attributes, func(attribute entities.Attribute) (string, entities.Attribute) {
		return attribute.Code, attribute
	})
}

func convertBookAttributes(attributes map[string]entities.Attribute, bookAttributes map[string][]string) []entities.AttributeToWeb {
	result := pkg.MapToSlice(bookAttributes, func(code string, values []string) entities.AttributeToWeb {
		return entities.AttributeToWeb{
			Code:   code,
			Name:   attributes[code].PluralName,
			Values: values,
		}
	})

	slices.SortStableFunc(result, func(a, b entities.AttributeToWeb) int {
		return attributes[a.Code].Order - attributes[b.Code].Order
	})

	return result
}
