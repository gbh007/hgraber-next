package bookrequester

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	BookIDs(ctx context.Context, filter entities.BookFilter) ([]uuid.UUID, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
	BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookPages(ctx context.Context, bookID uuid.UUID) ([]entities.Page, error)
	Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error)
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

func (uc *UseCase) Books(ctx context.Context, filter entities.BookFilter) ([]entities.BookFull, error) {
	ids, err := uc.storage.BookIDs(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get ids :%w", err)
	}

	books := make([]entities.BookFull, len(ids))

	for i, id := range ids {
		book, err := uc.requestBook(ctx, bookRequest{
			ID:                id,
			IncludeAttributes: true,
			IncludePages:      true,
			IncludeLabels:     true,
		})
		if err != nil {
			return nil, fmt.Errorf("book %s :%w", id.String(), err)
		}

		books[i] = book
	}

	return books, nil
}

func (uc *UseCase) Book(ctx context.Context, bookID uuid.UUID) (entities.Book, error) {
	book, err := uc.requestBook(ctx, bookRequest{
		ID: bookID,
	})
	if err != nil {
		return entities.Book{}, err
	}

	return book.Book, nil
}

func (uc *UseCase) BookFull(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error) {
	book, err := uc.requestBook(ctx, bookRequest{
		ID:                bookID,
		IncludeAttributes: true,
		IncludePages:      true,
		IncludeLabels:     true,
	})
	if err != nil {
		return entities.BookFull{}, err
	}

	return book, nil
}
