package book

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) SetBookRebuild(ctx context.Context, bookID uuid.UUID, reBuilded bool) error {
	res, err := repo.Pool.Exec(
		ctx,
		`UPDATE books SET is_rebuild = $2 WHERE id = $1;`,
		bookID, reBuilded,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.BookNotFoundError
	}

	return nil
}
