package parsing

import (
	"context"

	"hgnext/internal/entities"
)

type Storage interface {
	NewBook(ctx context.Context, book entities.Book) error
}
