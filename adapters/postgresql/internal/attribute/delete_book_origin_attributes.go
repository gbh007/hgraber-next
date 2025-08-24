package attribute

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *AttributeRepo) DeleteBookOriginAttributes(ctx context.Context, bookID uuid.UUID) error {
	_, err := repo.Pool.Exec(ctx, `DELETE FROM book_origin_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
