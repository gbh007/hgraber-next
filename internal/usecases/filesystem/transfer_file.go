package filesystem

import (
	"context"
	"fmt"

	"hgnext/internal/entities"
)

func (uc *UseCase) TransferFile(ctx context.Context, transfer entities.FileTransfer) error {
	file, err := uc.storage.File(ctx, transfer.FileID)
	if err != nil {
		return fmt.Errorf("storage: get file: %w", err)
	}

	if file.FSID == transfer.FSID {
		return nil
	}

	body, err := uc.fileStorage.Get(ctx, file.ID, &file.FSID)
	if err != nil {
		return fmt.Errorf("file storage: get file body: %w", err)
	}

	err = uc.fileStorage.Create(ctx, file.ID, body, transfer.FSID)
	if err != nil {
		return fmt.Errorf("file storage: create file body: %w", err)
	}

	// Перед обновлением данных в БД валидируем новые данные.
	err = uc.validateFileBody(ctx, file.ID, file.Hash(), transfer.FSID)
	if err != nil {
		return fmt.Errorf("validate body: %w", err)
	}

	err = uc.storage.UpdateFileFS(ctx, file.ID, transfer.FSID)
	if err != nil {
		return fmt.Errorf("storage: update file fs: %w", err)
	}

	// Удаляем данные из старой ФС
	err = uc.fileStorage.Delete(ctx, file.ID, &file.FSID)
	if err != nil {
		return fmt.Errorf("file storage: delete file body: %w", err)
	}

	return nil
}

func (uc *UseCase) FileTransferList() []entities.FileTransfer {
	return uc.tmpStorage.FileTransferList()
}
