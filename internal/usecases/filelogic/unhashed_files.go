package filelogic

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"

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

	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("read body for hashing: %w", err)
	}

	f.Size = int64(len(data))
	f.Md5Sum = fmt.Sprintf("%x", md5.Sum(data))
	f.Sha256Sum = fmt.Sprintf("%x", sha256.Sum256(data))

	err = uc.storage.UpdateFileHash(ctx, f.ID, f.Md5Sum, f.Sha256Sum, f.Size)
	if err != nil {
		return fmt.Errorf("update hash info in storage: %w", err)
	}

	return nil
}
