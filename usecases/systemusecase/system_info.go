package systemusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func (uc *UseCase) SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error) {
	systemSize, err := uc.storage.SystemSize(ctx)
	if err != nil {
		return systemmodel.SystemSizeInfo{}, fmt.Errorf("storage: %w", err)
	}

	return systemSize, nil
}

func (uc *UseCase) WorkersInfo(ctx context.Context) []systemmodel.SystemWorkerStat {
	return uc.workerManager.Info()
}
