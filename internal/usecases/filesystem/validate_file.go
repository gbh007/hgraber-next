package filesystem

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) ValidateFS(ctx context.Context, fsID uuid.UUID) error {
	ids, err := uc.storage.FileIDsByFS(ctx, fsID)
	if err != nil {
		return err
	}

	uc.tmpStorage.AddToValidate(ids)

	return nil
}

func (uc *UseCase) ValidateFile(ctx context.Context, fileID uuid.UUID) error {
	file, err := uc.storage.File(ctx, fileID)
	if err != nil {
		return fmt.Errorf("storage: get file: %w", err)
	}

	body, err := uc.fileStorage.Get(ctx, file.ID, &file.FSID)

	if errors.Is(err, entities.FileNotFoundError) {
		uc.logger.DebugContext(
			ctx, "missing file in fs",
			slog.String("id", file.ID.String()),
		)

		err = uc.storage.UpdateFileInvalidData(ctx, file.ID, true)
		if err != nil {
			return fmt.Errorf("storage: update invalid file data: missing: %w", err)
		}
	}

	if err != nil {
		return fmt.Errorf("file storage: get file body: %w", err)
	}

	newHash, err := entities.HashFile(body)
	if err != nil {
		return fmt.Errorf("hash file: %w", err)
	}

	invalidData := file.Hash() != newHash

	err = uc.storage.UpdateFileInvalidData(ctx, file.ID, invalidData)
	if err != nil {
		return fmt.Errorf("storage: update invalid file data: %w", err)
	}

	return nil
}

func (uc *UseCase) FileIDsToValidate() []uuid.UUID {
	return uc.tmpStorage.ValidateList()
}
