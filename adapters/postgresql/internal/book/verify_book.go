package book

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error {
	res, err := repo.Pool.Exec(
		ctx,
		`UPDATE books SET verified_at = $2, verified = $3 WHERE id = $1;`,
		bookID, model.TimeToDB(verifiedAt), verified,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrBookNotFound
	}

	return nil
}
