package webapi

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/entities"
)

func (uc *UseCase) SystemSize(ctx context.Context) (entities.SystemSizeInfo, error) {
	systemSize, err := uc.storage.SystemSize(ctx)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("storage: %w", err)
	}

	return systemSize, nil
}

func (uc *UseCase) WorkersInfo(ctx context.Context) []entities.SystemWorkerStat {
	return uc.workerManager.Info()
}
