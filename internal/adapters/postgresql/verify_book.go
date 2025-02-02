package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/internal/entities"
)

func (d *Database) VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error {
	res, err := d.db.ExecContext(
		ctx,
		`UPDATE books SET verified_at = $2, verified = $3 WHERE id = $1;`,
		bookID.String(), model.TimeToDB(verifiedAt), verified,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if !d.isApply(ctx, res) {
		return entities.BookNotFoundError
	}

	return nil
}
