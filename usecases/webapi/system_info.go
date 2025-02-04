package webapi

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) SystemSize(ctx context.Context) (core.SystemSizeInfo, error) {
	systemSize, err := uc.storage.SystemSize(ctx)
	if err != nil {
		return core.SystemSizeInfo{}, fmt.Errorf("storage: %w", err)
	}

	return systemSize, nil
}

func (uc *UseCase) WorkersInfo(ctx context.Context) []core.SystemWorkerStat {
	return uc.workerManager.Info()
}
