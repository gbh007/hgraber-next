package webapi

import (
	"context"
	"fmt"

	"hgnext/internal/entities"
)

func (uc *UseCase) SystemInfo(ctx context.Context) (entities.SystemSizeInfoWithMonitor, error) {
	systemSize, err := uc.storage.SystemSize(ctx)
	if err != nil {
		return entities.SystemSizeInfoWithMonitor{}, fmt.Errorf("storage: %w", err)
	}

	workers := uc.workerManager.Info()

	return entities.SystemSizeInfoWithMonitor{
		SystemSizeInfo: systemSize,
		Workers:        workers,
	}, nil
}
