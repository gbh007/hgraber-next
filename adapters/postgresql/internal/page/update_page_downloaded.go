package page

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) UpdatePageDownloaded(
	ctx context.Context,
	id uuid.UUID,
	pageNumber int,
	downloaded bool,
	fileID uuid.UUID,
) error {
	res, err := repo.Pool.Exec(
		ctx,
		`UPDATE pages SET downloaded = $1, load_at = $2, file_id = $5 WHERE book_id = $3 AND page_number = $4;`,
		downloaded, time.Now().UTC(), id, pageNumber, model.UUIDToDB(fileID),
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrPageNotFound
	}

	return nil
}
