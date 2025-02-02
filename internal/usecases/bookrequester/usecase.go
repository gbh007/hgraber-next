package bookrequester

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

type storage interface {
	BookIDs(ctx context.Context, filter entities.BookFilter) ([]uuid.UUID, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
	BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookPages(ctx context.Context, bookID uuid.UUID) ([]entities.Page, error)
	Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error)

	BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error)
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.DeadHash, error)
}

type UseCase struct {
	logger *slog.Logger

	storage storage
}

func New(
	logger *slog.Logger,
	storage storage,
) *UseCase {
	return &UseCase{
		logger:  logger,
		storage: storage,
	}
}

func (uc *UseCase) Books(ctx context.Context, filter entities.BookFilter) ([]entities.BookContainer, error) {
	ids, err := uc.storage.BookIDs(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get ids :%w", err)
	}

	books := make([]entities.BookContainer, len(ids))

	for i, id := range ids {
		book, err := uc.requestBook(ctx, bookRequest{
			ID:                      id,
			IncludeOriginAttributes: filter.OriginAttributes,
			IncludeAttributes:       !filter.OriginAttributes,
			IncludePages:            true,
			IncludeLabels:           true,
		})
		if err != nil {
			return nil, fmt.Errorf("book %s :%w", id.String(), err)
		}

		books[i] = book
	}

	return books, nil
}

func (uc *UseCase) BookOriginFull(ctx context.Context, bookID uuid.UUID) (entities.BookContainer, error) {
	book, err := uc.requestBook(ctx, bookRequest{
		ID:                      bookID,
		IncludeOriginAttributes: true,
		IncludePages:            true,
		IncludeLabels:           true,
	})
	if err != nil {
		return entities.BookContainer{}, err
	}

	return book, nil
}
