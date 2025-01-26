package filesystem

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	FileStorage(ctx context.Context, id uuid.UUID) (entities.FileStorageSystem, error)
	NewFileStorage(ctx context.Context, fs entities.FileStorageSystem) error
	UpdateFileStorage(ctx context.Context, fs entities.FileStorageSystem) error
	DeleteFileStorage(ctx context.Context, id uuid.UUID) error

	FSFilesInfo(ctx context.Context, fsID uuid.UUID, onlyInvalidData bool) (entities.FSFilesInfo, error)
}

type fileStorage interface {
	FSList(ctx context.Context) ([]entities.FSWithStatus, error)
	FSChange(ctx context.Context, fsID uuid.UUID, deleted bool) error
}

type UseCase struct {
	storage     storage
	fileStorage fileStorage
}

func New(
	storage storage,
	fileStorage fileStorage,
) *UseCase {
	return &UseCase{
		storage:     storage,
		fileStorage: fileStorage,
	}
}

func (uc *UseCase) FileStoragesWithStatus(ctx context.Context) ([]entities.FSWithStatus, error) {
	storages, err := uc.fileStorage.FSList(ctx)
	if err != nil {
		return nil, fmt.Errorf("file storage: get fs list: %w", err)
	}

	for i, storage := range storages {
		info, err := uc.storage.FSFilesInfo(ctx, storage.Info.ID, false)
		if err != nil {
			return nil, fmt.Errorf("storage: get files info (%s): %w", storage.Info.ID.String(), err)
		}

		storages[i].DBFile = info
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
