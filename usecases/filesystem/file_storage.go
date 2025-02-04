package filesystem

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/entities"
)

func (uc *UseCase) FileStoragesWithStatus(ctx context.Context, includeDBInfo, includeAvailableSizeInfo bool) ([]entities.FSWithStatus, error) {
	storages, err := uc.fileStorage.FSList(ctx)
	if err != nil {
		return nil, fmt.Errorf("file storage: get fs list: %w", err)
	}

	if includeDBInfo || includeAvailableSizeInfo {
		for i, storage := range storages {
			if includeDBInfo {
				info, err := uc.storage.FSFilesInfo(ctx, storage.Info.ID, false, false)
				if err != nil {
					return nil, fmt.Errorf("storage: get files info (%s): %w", storage.Info.ID.String(), err)
				}

				storages[i].DBFile = &info

				invalidInfo, err := uc.storage.FSFilesInfo(ctx, storage.Info.ID, true, false)
				if err != nil {
					return nil, fmt.Errorf("storage: get invalid files info (%s): %w", storage.Info.ID.String(), err)
				}

				storages[i].DBInvalidFile = &invalidInfo

				detachedInfo, err := uc.storage.FSFilesInfo(ctx, storage.Info.ID, false, true)
				if err != nil {
					return nil, fmt.Errorf("storage: get detached files info (%s): %w", storage.Info.ID.String(), err)
				}

				storages[i].DBDetachedFile = &detachedInfo
			}

			if includeAvailableSizeInfo {
				state, err := uc.fileStorage.State(ctx, false, false, storage.Info.ID)
				if err != nil {
					return nil, fmt.Errorf("file storage: get state (%s): %w", storage.Info.ID.String(), err)
				}

				storages[i].AvailableSize = state.AvailableSize
			}
		}
	}

	slices.SortStableFunc(storages, func(a, b entities.FSWithStatus) int {
		return a.Info.CreatedAt.Compare(b.Info.CreatedAt)
	})

	return storages, nil
}

func (uc *UseCase) FileStorage(ctx context.Context, id uuid.UUID) (entities.FileStorageSystem, error) {
	return uc.storage.FileStorage(ctx, id)
}

func (uc *UseCase) NewFileStorage(ctx context.Context, fs entities.FileStorageSystem) (uuid.UUID, error) {
	fs.ID = uuid.Must(uuid.NewV7())
	fs.CreatedAt = time.Now().UTC()

	err := uc.storage.NewFileStorage(ctx, fs)
	if err != nil {
		return uuid.Nil, fmt.Errorf("storage: %w", err)
	}

	err = uc.fileStorage.FSChange(ctx, fs.ID, fs.NotAvailable())
	if err != nil {
		return uuid.Nil, fmt.Errorf("file storage: change fs: %w", err)
	}

	return fs.ID, nil
}

func (uc *UseCase) UpdateFileStorage(ctx context.Context, fs entities.FileStorageSystem) error {
	err := uc.storage.UpdateFileStorage(ctx, fs)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	err = uc.fileStorage.FSChange(ctx, fs.ID, fs.NotAvailable())
	if err != nil {
		return fmt.Errorf("file storage: change fs: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteFileStorage(ctx context.Context, id uuid.UUID) error {
	err := uc.storage.DeleteFileStorage(ctx, id)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	err = uc.fileStorage.FSChange(ctx, id, true)
	if err != nil {
		return fmt.Errorf("file storage: change fs: %w", err)
	}

	return nil
}
