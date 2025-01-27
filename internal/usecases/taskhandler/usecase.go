package taskhandler

import (
	"context"
	"log/slog"

	"hgnext/internal/entities"
)

type storage interface {
	SaveTask(task entities.RunnableTask)
	GetTaskResults() []*entities.TaskResult
}

type deduplicator interface {
	DeduplicateFiles(ctx context.Context) (entities.RunnableTask, error)
	FillDeadHashes(ctx context.Context, withRemoveDeletedPages bool) (entities.RunnableTask, error)
}

type cleanuper interface {
	RemoveDetachedFiles(ctx context.Context) (entities.RunnableTask, error)
	RemoveFilesInStoragesMismatch(ctx context.Context) (entities.RunnableTask, error)
	CleanDeletedPages(ctx context.Context) (entities.RunnableTask, error)
	CleanDeletedRebuilds(ctx context.Context) (entities.RunnableTask, error)
}

type UseCase struct {
	logger *slog.Logger

	storage      storage
	deduplicator deduplicator
	cleanuper    cleanuper
}

func New(
	logger *slog.Logger,
	storage storage,
	deduplicator deduplicator,
	cleanuper cleanuper,
) *UseCase {
	return &UseCase{
		logger:       logger,
		storage:      storage,
		deduplicator: deduplicator,
		cleanuper:    cleanuper,
	}
}

func (uc *UseCase) RunTask(ctx context.Context, code entities.TaskCode) error {
	var (
		task entities.RunnableTask
		err  error
	)

	switch code {
	case entities.DeduplicateFilesTaskCode:
		task, err = uc.deduplicator.DeduplicateFiles(ctx)
	case entities.RemoveDetachedFilesTaskCode:
		task, err = uc.cleanuper.RemoveDetachedFiles(ctx)
	case entities.RemoveFilesInStoragesMismatchTaskCode:
		task, err = uc.cleanuper.RemoveFilesInStoragesMismatch(ctx)
	case entities.FillDeadHashesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, false)
	case entities.FillDeadHashesAndRemoveDeletedPagesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, true)
	case entities.CleanDeletedPagesTaskCode:
		task, err = uc.cleanuper.CleanDeletedPages(ctx)
	case entities.CleanDeletedRebuildsTaskCode:
		task, err = uc.cleanuper.CleanDeletedRebuilds(ctx)
	}

	if err != nil {
		return err
	}

	if task != nil {
		uc.storage.SaveTask(task)
	}

	return nil
}

func (uc *UseCase) TaskResults(ctx context.Context) ([]*entities.TaskResult, error) {
	return uc.storage.GetTaskResults(), nil
}
