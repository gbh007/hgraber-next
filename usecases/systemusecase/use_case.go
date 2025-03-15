package systemusecase

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type storage interface {
	SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error)
}

type tmpStorage interface {
	SaveTask(task systemmodel.RunnableTask)
	GetTaskResults() []*systemmodel.TaskResult
}

type deduplicator interface {
	DeduplicateFiles(ctx context.Context) (systemmodel.RunnableTask, error)
	FillDeadHashes(ctx context.Context, withRemoveDeletedPages bool) (systemmodel.RunnableTask, error)
}

type cleanuper interface {
	RemoveDetachedFiles(ctx context.Context) (systemmodel.RunnableTask, error)
	RemoveFilesInStoragesMismatch(ctx context.Context, fsID uuid.UUID) (systemmodel.RunnableTask, error)
	CleanDeletedPages(ctx context.Context) (systemmodel.RunnableTask, error)
	CleanDeletedRebuilds(ctx context.Context) (systemmodel.RunnableTask, error)
}

type workerManager interface {
	Info() []systemmodel.SystemWorkerStat

	SetRunnerCount(ctx context.Context, counts map[string]int)
}

type attributeRemaper interface {
	RemapBooks(ctx context.Context) (systemmodel.RunnableTask, error)
}

type UseCase struct {
	logger *slog.Logger

	storage          storage
	tmpStorage       tmpStorage
	deduplicator     deduplicator
	cleanuper        cleanuper
	workerManager    workerManager
	attributeRemaper attributeRemaper
}

func New(
	logger *slog.Logger,
	storage storage,
	tmpStorage tmpStorage,
	deduplicator deduplicator,
	cleanuper cleanuper,
	workerManager workerManager,
	attributeRemaper attributeRemaper,
) *UseCase {
	return &UseCase{
		logger:           logger,
		storage:          storage,
		tmpStorage:       tmpStorage,
		deduplicator:     deduplicator,
		cleanuper:        cleanuper,
		workerManager:    workerManager,
		attributeRemaper: attributeRemaper,
	}
}
