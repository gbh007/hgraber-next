package taskhandler

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type storage interface {
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

func (uc *UseCase) RunTask(ctx context.Context, code systemmodel.TaskCode) error {
	var (
		task systemmodel.RunnableTask
		err  error
	)

	switch code {
	case systemmodel.DeduplicateFilesTaskCode:
		task, err = uc.deduplicator.DeduplicateFiles(ctx)
	case systemmodel.RemoveDetachedFilesTaskCode:
		task, err = uc.cleanuper.RemoveDetachedFiles(ctx)
	case systemmodel.FillDeadHashesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, false)
	case systemmodel.FillDeadHashesAndRemoveDeletedPagesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, true)
	case systemmodel.CleanDeletedPagesTaskCode:
		task, err = uc.cleanuper.CleanDeletedPages(ctx)
	case systemmodel.CleanDeletedRebuildsTaskCode:
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

func (uc *UseCase) TaskResults(ctx context.Context) ([]*systemmodel.TaskResult, error) {
	return uc.storage.GetTaskResults(), nil
}

func (uc *UseCase) RemoveFilesInFSMismatch(ctx context.Context, fsID uuid.UUID) error {
	task, err := uc.cleanuper.RemoveFilesInStoragesMismatch(ctx, fsID)
	if err != nil {
		return err
	}

	uc.storage.SaveTask(task)

	return nil
}
