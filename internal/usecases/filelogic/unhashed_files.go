package filelogic

import (
	"context"
	"fmt"

	"hgnext/internal/entities"
)

func (uc *UseCase) UnHashedFiles(ctx context.Context) ([]entities.File, error) {
	return uc.storage.GetUnHashedFiles(ctx)
}

func (uc *UseCase) HandleFileHash(ctx context.Context, f entities.File) error {
	body, err := uc.fileStorage.Get(ctx, f.ID)
	if err != nil {
		return fmt.Errorf("get file body: %w", err)
	}

	hash, err := entities.HashFile(body)
	if err != nil {
		return fmt.Errorf("hash file: %w", err)
	}

	f.Size = hash.Size
	f.Md5Sum = hash.Md5Sum
	f.Sha256Sum = hash.Sha256Sum

	err = uc.storage.UpdateFileHash(ctx, f.ID, f.Md5Sum, f.Sha256Sum, f.Size)
	if err != nil {
		return fmt.Errorf("update hash info in storage: %w", err)
	}

	return nil
}
