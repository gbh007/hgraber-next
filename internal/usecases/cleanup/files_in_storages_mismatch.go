package cleanup

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/pkg"
)

func (uc *UseCase) RemoveFilesInStoragesMismatch(ctx context.Context) (notInDBCount, notInFSCount int, err error) {
	fileIDs, err := uc.fileStorage.IDs(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("file storage get ids: %w", err)
	}

	storageIDs, err := uc.storage.FileIDs(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("storage get ids: %w", err)
	}

	fileNotInDB := pkg.SliceToMap(fileIDs, func(id uuid.UUID) (uuid.UUID, struct{}) {
		return id, struct{}{}
	})
	fileNotInFS := pkg.SliceToMap(storageIDs, func(id uuid.UUID) (uuid.UUID, struct{}) {
		return id, struct{}{}
	})

	for _, id := range fileIDs {
		delete(fileNotInFS, id)
	}

	for _, id := range storageIDs {
		delete(fileNotInDB, id)
	}

	uc.logger.DebugContext(
		ctx, "RemoveFilesInStoragesMismatch",
		slog.Int("file_not_in_fs_count", len(fileNotInFS)),
		slog.Int("file_not_in_db_count", len(fileNotInDB)),
		slog.Any("file_not_in_fs", fileNotInFS),
		slog.Any("file_not_in_db", fileNotInDB),
	)

	for id := range fileNotInDB {
		err = uc.fileStorage.Delete(ctx, id)
		if err != nil {
			return 0, 0, fmt.Errorf("file storage remove (%s): %w", id.String(), err)
		}
	}

	for id := range fileNotInFS {
		err = uc.storage.DeleteFile(ctx, id)
		if err != nil {
			return 0, 0, fmt.Errorf("storage remove (%s): %w", id.String(), err)
		}
	}

	return len(fileNotInDB), len(fileNotInFS), nil
}
