package systemusecase

import "context"

func (uc *UseCase) SetWorkerConfig(ctx context.Context, counts map[string]int) {
	uc.workerManager.SetRunnerCount(ctx, counts)
}
