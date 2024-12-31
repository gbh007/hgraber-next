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
	IncludeLabels           bool
	IncludeSize             bool
}

func (uc *UseCase) requestBook(ctx context.Context, req bookRequest) (entities.BookFull, error) {
	b, err := uc.storage.GetBook(ctx, req.ID)
	if err != nil {
		return entities.BookFull{}, fmt.Errorf("get book :%w", err)
	}

	out := entities.BookFull{
		Book: b,
	}

	switch {
	case req.IncludeOriginAttributes:
		attributes, err := uc.storage.BookOriginAttributes(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get attributes :%w", err)
		}

		out.Attributes = attributes

	case req.IncludeAttributes:
		attributes, err := uc.storage.BookAttributes(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get attributes :%w", err)
		}

		out.Attributes = attributes
	}

	if req.IncludePages {
		pages, err := uc.storage.BookPages(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get pages :%w", err)
		}

		out.Pages = pages
	}

	if req.IncludeLabels {
		labels, err := uc.storage.Labels(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get labels :%w", err)
		}

		out.Labels = labels
	}

	if req.IncludeSize {
		size, err := uc.BookSize(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get size :%w", err)
		}

		out.Size = size
	}

	return out, nil
}

func (uc *UseCase) BookSize(ctx context.Context, originBookID uuid.UUID) (entities.BookSize, error) {
	bookPages, err := uc.storage.BookPagesWithHash(ctx, originBookID)
	if err != nil {
		return entities.BookSize{}, fmt.Errorf("get book hashes storage: %w", err)
	}

	fileCounts := make(map[entities.FileHash]int, len(bookPages))

	md5Sums := make([]string, len(bookPages))
	for i, page := range bookPages {
		md5Sums[i] = page.Md5Sum

		fileCounts[page.Hash()] = 1 // Это условие (значение 1) нужно, чтобы дубликаты внутри книги не дали ложно-положительного срабатывания
	}

	bookIDs, err := uc.storage.BookIDsByMD5(ctx, md5Sums)
	if err != nil {
		return entities.BookSize{}, fmt.Errorf("get books by md5 from storage: %w", err)
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
			return entities.BookSize{}, fmt.Errorf("get pages (%s) from storage: %w", bookID.String(), err)
		}

		for _, page := range pages {
			if _, ok := fileCounts[page.Hash()]; ok {
				fileCounts[page.Hash()]++
			}
		}
	}

	result := entities.BookSize{}

	for _, page := range bookPages {
		if c, ok := fileCounts[page.Hash()]; ok {
			if c > 1 {
				result.Shared += page.Size
			} else {
				result.Unique += page.Size
			}

			delete(fileCounts, page.Hash()) // Это нужно, чтобы дубликаты внутри книги не увеличивали уникальный объем
		}

		result.Total += page.Size
	}

	return result, nil
}
