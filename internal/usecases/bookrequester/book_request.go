package bookrequester

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type bookRequest struct {
	ID                      uuid.UUID
	IncludeAttributes       bool
	IncludeOriginAttributes bool
	IncludePages            bool
	IncludePagesWithHash    bool
	IncludeLabels           bool
	IncludeSize             bool
}

// FIXME: избавится от этого непотребства, вынести в отдельные методы получение списка книг и получение детальных данных по 1 книге.
func (uc *UseCase) requestBook(ctx context.Context, req bookRequest) (entities.BookContainer, error) {
	b, err := uc.storage.GetBook(ctx, req.ID)
	if err != nil {
		return entities.BookContainer{}, fmt.Errorf("get book: %w", err)
	}

	out := entities.BookContainer{
		Book: b,
	}

	switch {
	case req.IncludeOriginAttributes:
		attributes, err := uc.storage.BookOriginAttributes(ctx, req.ID)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get attributes: %w", err)
		}

		out.Attributes = attributes

	case req.IncludeAttributes:
		attributes, err := uc.storage.BookAttributes(ctx, req.ID)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get attributes: %w", err)
		}

		out.Attributes = attributes
	}

	if req.IncludePagesWithHash {
		pages, err := uc.storage.BookPagesWithHash(ctx, req.ID)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get pages with hash:%w", err)
		}

		out.PagesWithHash = pages

		if req.IncludePages {
			out.Pages = make([]entities.Page, 0, len(pages))
			for _, page := range pages {
				out.Pages = append(out.Pages, page.Page)
			}
		}
	} else if req.IncludePages {
		pages, err := uc.storage.BookPages(ctx, req.ID)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get pages: %w", err)
		}

		out.Pages = pages
	}

	if req.IncludeLabels {
		labels, err := uc.storage.Labels(ctx, req.ID)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get labels: %w", err)
		}

		out.Labels = labels
	}

	if req.IncludeSize {
		size, deadHashOnPage, err := uc.BookSize(ctx, req.ID)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get size: %w", err)
		}

		out.Size = size
		out.DeadHashOnPage = deadHashOnPage
	}

	return out, nil
}

// FIXME: крайне тяжелые данные, их необходимо вынести отдельно (и в юзкейсы дедупливатора)
func (uc *UseCase) BookSize(ctx context.Context, originBookID uuid.UUID) (entities.BookSize, map[int]struct{}, error) {
	bookPages, err := uc.storage.BookPagesWithHash(ctx, originBookID)
	if err != nil {
		return entities.BookSize{}, nil, fmt.Errorf("get book hashes storage: %w", err)
	}

	fileCounts := make(map[entities.FileHash]int, len(bookPages))

	md5Sums := make([]string, len(bookPages))
	for i, page := range bookPages {
		md5Sums[i] = page.Md5Sum

		fileCounts[page.FileHash] = 1 // Это условие (значение 1) нужно, чтобы дубликаты внутри книги не дали ложно-положительного срабатывания
	}

	bookIDs, err := uc.storage.BookIDsByMD5(ctx, md5Sums)
	if err != nil {
		return entities.BookSize{}, nil, fmt.Errorf("get books by md5 from storage: %w", err)
	}

	bookHandled := make(map[uuid.UUID]struct{}, len(bookIDs))
	bookHandled[originBookID] = struct{}{}

	for _, bookID := range bookIDs {
		if _, ok := bookHandled[bookID]; ok {
			continue
		}

		bookHandled[bookID] = struct{}{}

		pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
		if err != nil {
			return entities.BookSize{}, nil, fmt.Errorf("get pages (%s) from storage: %w", bookID.String(), err)
		}

		for _, page := range pages {
			if _, ok := fileCounts[page.FileHash]; ok {
				fileCounts[page.FileHash]++
			}
		}
	}

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return entities.BookSize{}, nil, fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[entities.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	result := entities.BookSize{}

	deadHashOnPage := make(map[int]struct{}, len(bookPages))

	for _, page := range bookPages {
		_, hasDeadHash := existsDeadHashes[page.FileHash]
		if hasDeadHash {
			deadHashOnPage[page.PageNumber] = struct{}{}
		}

		if c, ok := fileCounts[page.FileHash]; ok {
			if c > 1 {
				result.Shared += page.Size
				result.SharedCount++
			} else {
				result.Unique += page.Size
				result.UniqueCount++

				if !hasDeadHash {
					result.UniqueWithoutDeadHashes += page.Size
					result.UniqueWithoutDeadHashesCount++
				}
			}

			if hasDeadHash {
				result.DeadHashes += page.Size
				result.DeadHashesCount++
			}

			delete(fileCounts, page.FileHash) // Это нужно, чтобы дубликаты внутри книги не увеличивали уникальный объем
		} else {
			result.InnerDuplicateCount++
		}

		result.Total += page.Size
	}

	return result, deadHashOnPage, nil
}
