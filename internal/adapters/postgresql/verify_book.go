package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (d *Database) VerifyBook(ctx context.Context, bookID uuid.UUID) error {
	res, err := d.db.ExecContext(
		ctx,
		`UPDATE books SET verified_at = $2, verified = $3 WHERE id = $1;`,
		bookID.String(), time.Now().UTC(), true,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if !d.isApply(ctx, res) {
		return entities.BookNotFoundError
	}

	return nil
}
